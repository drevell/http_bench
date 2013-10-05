This is a toy program that's intended to show that passing protobufs over
HTTP 1.1 using Go is a reasonably fast way to do RPC.

This is just a simple throughput/latency test of a single-threaded HTTP
request/response loop. On each iteration of the loop, the client sends a 
"ping" protobuf inside an HTTP request and expects a "pong" protobuf response.

I found mean latency of 0.5ms (498us) per request/response. This corresponds to 
single-connection throughput of 2006 reqs/sec. In this test, both the client and
server were t1.micro instances in the same AWS availability zone 
(us-west-2) running ubuntu 13.04 and go 1.1.2.

The performance seems to rely heavily on the connection pooling that's built in
to go's net/http package. The first request seems to take about 2ms, and that's
probably due to TCP connection establishment overhead.

Go seems to disable Nagle's Algorithm for TCP connections by default (see
http://golang.org/pkg/net/#TCPConn.SetNoDelay ). This probably has the effect
of increasing performance for this kind of small request/response workload.

Future work:

 - Measure the network latency to determine protocol overhead as a fraction
