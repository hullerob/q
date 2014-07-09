// © 2014, Robert Hülle

package main

import "log"

func main() {
	queue := NewQueue()
	end := make(chan bool)
	queue.Run()
	qrpc, err := startRpc(queue, end)
	if err != nil {
		log.Fatalf("Could not start qd: %v", err)
	}
	<-end
	err = qrpc.stopRPC()
	if err != nil {
		log.Fatalf("Could not stop qd: %v", err)
	}
}
