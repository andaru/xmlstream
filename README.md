[![travis-ci](https://travis-ci.org/andaru/xmlstream.svg?branch=master)](https://travis-ci.org/andaru/xmlstream)
[![GoDoc](https://godoc.org/github.com/andaru/xmlstream?status.svg)](https://godoc.org/github.com/andaru/xmlstream)

# xmlstream #

The `xmlstream` package provides Golang libraries for building event
driven XML stream processors, such as
[XMPP](https://tools.ietf.org/html/rfc3920) and
[NETCONF](https://tools.ietf.org/html/rfc6241) stream processing.

Schema are constructed as a tree of `*xmlstream.Node`, and a state
machine `xmlstream.StateMachine` is provided. A parser object
combining the XML decoder with traversal of the schema tree via the
state machine is not included here, but can be found in a user
package, <https://github.com/andaru/netconf/schema>.
