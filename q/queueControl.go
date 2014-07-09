// © 2014, Robert Hülle

package main

func queueRunCommand(_ string, _ []string) error {
	err := doRPC("Qrpc.RunQueue", 0, nil)
	return err
}

func queueStopCommand(_ string, _ []string) error {
	err := doRPC("Qrpc.StopQueue", 0, nil)
	return err
}

func init() {
	rc := &clientCommand{
		f:    queueRunCommand,
		name: "R",
		args: "",
		help: "run queue",
	}
	registerCommand(rc)
	sc := &clientCommand{
		f:    queueStopCommand,
		name: "S",
		args: "",
		help: "stop queue",
	}
	registerCommand(sc)
}
