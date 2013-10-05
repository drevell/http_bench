This is a toy program that's intended to show that passing protobufs over
HTTP 1.1 using Go is a reasonably fast way to do RPC.

This is just a simple throughput/latency test of a single-threaded HTTP
request/response loop. On each iteration of the loop, the client sends a 
"ping" protobuf inside an HTTP request and expects a "pong" protobuf response.

I found mean latency of 0.5ms (498us) per request/response. This corresponds to 
single-connection throughput of 2006 reqs/sec. In this test, both the client and
server were t1.micro instances in the same AWS availability zone 
(us-west-2) running ubuntu 13.04 and go 1.1.2.

Future work:

 - Measure the network latency to determine protocol overhead as a fraction
