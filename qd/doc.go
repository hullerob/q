// © 2014, Robert Hülle

/*
Command `qd` runs daemon that listens for commands and executes them one by
one.

Qd listens on address defined in environment variable `QADDR`. Address format
is "<network>:<address>", where network and address are passed to `net.Listen`
and `rpc.Dial`.

Examples:

	QADDR="unix:/tmp/q.socket" qd

	QADDR="tcp:127.0.0.1:9876" qd

BUGS:

Qd does not clean socket file after termination.

*/
package main
