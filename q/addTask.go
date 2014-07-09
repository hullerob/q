// © 2014, Robert Hülle

package main

import (
	"errors"
	c "github.com/hullerob/q/common"
	"os"
	"os/exec"
)

func addTaskCommand(_ string, args []string) error {
	if len(args) == 0 {
		return errors.New("missing arguments")
	}
	cmd1 := exec.Command(args[0], args[1:]...)
	env := os.Environ()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd2 := &c.Command{
		Path: cmd1.Path,
		Args: cmd1.Args,
		Env:  env,
		Dir:  wd,
	}
	err = doRPC("Qrpc.AddTask", cmd2, nil)
	return err
}

func init() {
	ac := &clientCommand{
		f:    addTaskCommand,
		name: "a",
		args: "<command> [<args>] ...",
		help: "send new command to queue",
	}
	registerCommand(ac)
}
