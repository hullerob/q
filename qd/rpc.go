// © 2014, Robert Hülle

package main

import (
	c "github.com/hullerob/q/common"
	"net"
	"net/rpc"
	"os/exec"
)

type Qrpc struct {
	queue    *Queue
	end      chan<- bool
	listener net.Listener
}

func (q *Qrpc) AddTask(arg c.Command, reply *int) error {
	t := &Task{
		cmd:   exec.Command(arg.Path, arg.Args[1:]...),
		state: c.StReady,
		err:   nil,
	}
	t.cmd.Args[0] = arg.Args[0]
	t.cmd.Dir = arg.Dir
	t.cmd.Env = arg.Env
	q.queue.AddTask(t)
	return nil
}

func (q *Qrpc) GetQueueStatus(_ int, reply *c.QueueStatus) error {
	qStat := q.queue.QueueInfo()
	*reply = *qStat
	return nil
}

func (q *Qrpc) RunQueue(_ int, _ *int) error {
	q.queue.Run()
	return nil
}

func (q *Qrpc) StopQueue(_ int, _ *int) error {
	q.queue.Stop()
	return nil
}

func (q *Qrpc) stopRPC() error {
	return q.listener.Close()
}

func startRpc(queue *Queue, end chan<- bool) (*Qrpc, error) {
	listener, err := net.Listen(c.QNetwork, c.QLaddr)
	if err != nil {
		return nil, err
	}
	qrpc := &Qrpc{
		queue:    queue,
		end:      end,
		listener: listener,
	}
	server := rpc.NewServer()
	err = server.Register(qrpc)
	if err != nil {
		return nil, err
	}
	go server.Accept(listener)
	return qrpc, nil
}
