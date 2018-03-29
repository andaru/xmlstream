[![travis-ci](https://travis-ci.org/andaru/xmlstream.svg?branch=master)](https://travis-ci.org/andaru/xmlstream)
[![GoDoc](https://godoc.org/github.com/andaru/xmlstream?status.svg)](https://godoc.org/github.com/andaru/xmlstream)

# xmlstream #

The `xmlstream` package provides Golang libraries for building event
driven XML stream processors, such as
[XMPP](https://tools.ietf.org/html/rfc3920) and
[NETCONF](https://tools.ietf.org/html/rfc6241) stream processing.

Schema are constructed as a tree of `*xmlstream.Node`, and a state
machine `xmlstream.StateMachine` is provided.

See godoc for API documentation and an example showing how a
schema tree along with an XML decoder and the `StateMachine` combine
to parse and validate NETCONF server sessions.
