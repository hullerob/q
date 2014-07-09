// © 2014, Robert Hülle

package main

import (
	c "github.com/hullerob/q/common"
	"net/rpc"
)

func doRPC(procedure string, arg interface{}, reply interface{}) error {
	client, err := rpc.Dial(c.QNetwork, c.QLaddr)
	if err != nil {
		return err
	}
	err = client.Call(procedure, arg, reply)
	return err
}
