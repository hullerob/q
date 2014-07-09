Command queue
=============

About
-----

This is coding exercise inspired by [mkopta/q](https://github.com/mkopta/q).

This tool queues commands to execute them one by one.

Install
-------

    go get github.com/hullerob/q/{q,qd}

Example
-------

In one terminal:

    QADDR="unix:/tmp/q.sock" qd

In other terminal:

    export QADDR="unix:/tmp/q.sock"
    q a cp -r some/dir /mnt/some/slow/disk
    q a cp -r some/other/dir /mnt/some/slow/disk
    q l

Q Commands
---------

    q h          print help about commands
    q a [cmd]    add cmd to queue
    q i          print status of queue (length of queue)
    q l          print queued commands
    q r          print info about all commands
    q R          run queue
    q S          stop queue

There are currently no other commands.

What is QADDR
-------------

Anything in format `<network>:<addr>` that is accepted by Go's
`net.Listen(network, addr)` and `rpc.Dial(network, addr)`.

Examples:

* `unix:/tmp/socket`

* `tcp:127.0.0.1:9876`

BUGS
----

No support for clean qd shutdown, does not clean socket file.
