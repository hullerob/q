// © 2014, Robert Hülle

package main

import (
	c "github.com/hullerob/q/common"
	"os/exec"
)

type Task struct {
	cmd   *exec.Cmd
	id    int
	state c.TaskState
	err   error
}

func (t *Task) Run() error {
	return t.cmd.Run()
}

func (t *Task) Info() *c.TaskInfo {
	estr := ""
	if t.err != nil {
		estr = t.err.Error()
	}
	return &c.TaskInfo{
		Id:    t.id,
		Args:  t.cmd.Args,
		Wd:    t.cmd.Dir,
		State: t.state,
		Err:   estr,
	}
}
