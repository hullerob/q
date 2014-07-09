// © 2014, Robert Hülle

package main

import (
	"errors"
	"log"
	"os"
)

type clientCommand struct {
	f    func(string, []string) error
	name string
	args string
	help string
}

var (
	commands map[string]*clientCommand
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing arguments.")
	}
	cmd := os.Args[1]
	cmdFunc, e := commands[cmd]
	if !e {
		cmdFunc = commands["unknown"]
	}
	err := cmdFunc.f(os.Args[1], os.Args[2:])
	if err != nil {
		log.Fatalf("Command error: %v", err)
	}
}

func registerCommand(cmd *clientCommand) {
	if commands == nil {
		commands = make(map[string]*clientCommand)
	}
	_, e := commands[cmd.name]
	if e {
		log.Panicf("Registering one command twice: %s", cmd.name)
	}
	commands[cmd.name] = cmd
}

func unknownCommand(_ string, _ []string) error {
	log.Printf("Unknown command was called, try `%s h`", os.Args[0])
	return errors.New("unknown command")
}

func init() {
	uc := &clientCommand{
		f:    unknownCommand,
		name: "unknown",
		args: "",
		help: "unknown command was called",
	}
	registerCommand(uc)
}

func helpCommand(_ string, _ []string) error {
	log.Print("commands:")
	for name, cmd := range commands {
		log.Printf("%s %s \t%s", name, cmd.args, cmd.help)
	}
	return nil
}

func init() {
	hc := &clientCommand{
		f:    helpCommand,
		name: "h",
		args: "",
		help: "prints help message",
	}
	registerCommand(hc)
}
