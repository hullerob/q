// © 2014, Robert Hülle

package common

import (
	"bytes"
	"fmt"
)

const (
	MaxPrintArgs = 5
)

const (
	StReady TaskState = iota
	StRunning
	StDone
)

type TaskState int

func (ts TaskState) String() string {
	switch ts {
	case StReady:
		return "Q"
	case StRunning:
		return "R"
	case StDone:
		return "T"
	default:
		panic("TaskState: unexpected state")
	}
}

type QueueStatus struct {
	Stopped bool
	Waiting int
	Running int
	Done    int
	Tasks   []*TaskInfo
}

func (q *QueueStatus) String() string {
	var state string
	if q.Stopped {
		state = "stopped"
	} else {
		state = "running"
	}
	return fmt.Sprintf("queue is %s || waiting: %d | running: %d | done: %d",
		state, q.Waiting, q.Running, q.Done)
}

type Command struct {
	Path string
	Dir  string
	Args []string
	Env  []string
}

type TaskInfo struct {
	Id    int
	Args  []string
	Wd    string
	State TaskState
	Err   string
}

func (t TaskInfo) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "(%s)[%d]", t.State, t.Id)
	if t.State == StDone {
		if t.Err == "" {
			buf.WriteString("{0}")
		} else {
			fmt.Fprintf(buf, "{%s}", t.Err)
		}
	}
	fmt.Fprintf(buf, " wd: '%s' |", t.Wd)
	for i, a := range t.Args {
		if i >= MaxPrintArgs {
			buf.WriteString(" ...")
			break
		}
		buf.WriteString(" " + a)
	}
	return buf.String()
}
