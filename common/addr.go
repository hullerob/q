// © 2014, Robert Hülle

package common

import (
	"errors"
	"log"
	"os"
	"strings"
)

var (
	QNetwork string
	QLaddr   string
)

var (
	ErrorAddrNotSet    = errors.New("Environment variable QADDR not set.")
	ErrorAddrBadFormat = errors.New("Wrong format in QADDR.")
)

func init() {
	addr := os.Getenv("QADDR")
	var err error
	QNetwork, QLaddr, err = parseAddr(addr)
	if err != nil {
		log.Printf("warning: %v", err)
	}
}

func parseAddr(str string) (net string, addr string, err error) {
	if len(str) == 0 {
		err = ErrorAddrNotSet
		return
	}
	strs := strings.SplitN(str, ":", 2)
	if len(strs) != 2 {
		err = ErrorAddrBadFormat
		return
	}
	net = strs[0]
	addr = strs[1]
	return
}
