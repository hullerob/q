// © 2014, Robert Hülle

package main

import (
	c "github.com/hullerob/q/common"
	"log"
)

func queueStatusCommand(cmd string, _ []string) error {
	var status c.QueueStatus
	err := doRPC("Qrpc.GetQueueStatus", 0, &status)
	if err != nil {
		return err
	}
	log.Printf("Queue status at '%s:%s'", c.QNetwork, c.QLaddr)
	log.Printf("%s", &status)
	if cmd != "i" {
		for _, t := range status.Tasks {
			if cmd == "l" && t.State == c.StDone {
				continue
			}
			log.Printf("%s", t)
		}
	}
	return nil
}

func init() {
	qi := &clientCommand{
		f:    queueStatusCommand,
		name: "i",
		args: "",
		help: "print queue status",
	}
	registerCommand(qi)
	ql := &clientCommand{
		f:    queueStatusCommand,
		name: "l",
		args: "",
		help: "print queue list",
	}
	registerCommand(ql)
	qr := &clientCommand{
		f:    queueStatusCommand,
		name: "r",
		args: "",
		help: "print info about all tasks",
	}
	registerCommand(qr)
}
