// © 2014, Robert Hülle

package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parseAddressTests = []struct {
	in   string
	net  string
	addr string
	err  error
}{
	{"", "", "", ErrorAddrNotSet},
	{"foo", "", "", ErrorAddrBadFormat},
	{"tcp:127.0.0.1:9876", "tcp", "127.0.0.1:9876", nil},
	{"unix:/tmp/q.socket", "unix", "/tmp/q.socket", nil},
}

func TestParseAddress(t *testing.T) {
	for _, test := range parseAddressTests {
		net, addr, err := parseAddr(test.in)
		assert.Equal(t, test.net, net)
		assert.Equal(t, test.addr, addr)
		assert.Equal(t, test.err, err)
	}
}
