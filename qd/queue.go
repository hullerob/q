// © 2014, Robert Hülle

package main

import (
	c "github.com/hullerob/q/common"
	"sync"
)

const (
	QSStop = iota // Queue is stopped.
	QSRun         // Queue is running.
)

type Queue struct {
	state      int
	nextId     int
	queue      []*Task
	done       chan *Task
	exec       chan *Task
	wExec      chan *Task
	next       *Task
	loopTick   chan int
	loopEnd    chan int
	taskUpdate chan *c.TaskInfo
	mutex      sync.Mutex
	cWaiting   int
	cRunning   int
	cDone      int
	tinfo      map[int]*c.TaskInfo
	runCond    *sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{
		state:      QSStop,
		done:       make(chan *Task),
		loopTick:   make(chan int, 1),
		loopEnd:    make(chan int),
		taskUpdate: make(chan *c.TaskInfo),
		tinfo:      make(map[int]*c.TaskInfo),
		exec:       make(chan *Task),
		runCond:    sync.NewCond(new(sync.Mutex)),
	}
	go q.loop()
	go q.runner()
	return q
}

func (q *Queue) Run() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.state == QSRun {
		return
	}
	q.runCond.L.Lock()
	q.state = QSRun
	q.runCond.Broadcast()
	q.runCond.L.Unlock()
}

func (q *Queue) Stop() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.runCond.L.Lock()
	q.state = QSStop
	q.runCond.L.Unlock()
}

func (q *Queue) AddTask(nt *Task) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.append(nt)
	q.tick()
	return true
}

func (q *Queue) QueueInfo() *c.QueueStatus {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	ti := make([]*c.TaskInfo, len(q.tinfo))
	for k, v := range q.tinfo {
		ti[k] = v
	}
	qi := &c.QueueStatus{
		Stopped: q.state == QSStop,
		Waiting: q.cWaiting,
		Running: q.cRunning,
		Done:    q.cDone,
		Tasks:   ti,
	}
	return qi
}

func (q *Queue) runWait() {
	q.runCond.L.Lock()
	for q.state != QSRun {
		q.runCond.Wait()
	}
	q.runCond.L.Unlock()
}

func (q *Queue) loop() {
LOOP:
	q.runWait()
	q.runNext()
	select {
	case _ = <-q.loopTick:
	case _ = <-q.loopEnd:
		return
	case q.wExec <- q.next:
		q.taskRunning(q.next)
	case d := <-q.done:
		q.doneTask(d)
	}
	goto LOOP
}

func (q *Queue) taskRunning(t *Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.wExec = nil
	q.queue = q.queue[1:]
	q.next.state = c.StRunning
	q.cWaiting--
	q.cRunning++
	q.updateTaskInfo(q.next)
}

func (q *Queue) runNext() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.next != nil || len(q.queue) == 0 {
		return
	}
	q.next = q.queue[0]
	q.wExec = q.exec
}

func (q *Queue) tick() {
	select {
	case q.loopTick <- 1:
	default:
	}
}

func (q *Queue) append(nt *Task) {
	nt.id = q.nextId
	nt.state = c.StReady
	q.nextId++
	q.queue = append(q.queue, nt)
	q.cWaiting++
	q.updateTaskInfo(nt)
}

func (q *Queue) doneTask(t *Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	t.state = c.StDone
	q.next = nil
	q.cDone++
	q.cRunning--
	q.updateTaskInfo(t)
}

func (q *Queue) updateTaskInfo(t *Task) {
	q.tinfo[t.id] = t.Info()
}

func (q *Queue) runner() {
	for t := range q.exec {
		t.err = t.cmd.Run()
		q.done <- t
	}
	q.loopEnd <- 1
}
