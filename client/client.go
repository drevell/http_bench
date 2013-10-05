// A painfully simple latency benchmark for a simple protobuf-over-HTTP service.
// This is intended to show the best case latency for such a service.
package main

import (
	pbs ".."
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	numReq, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	targetUrl := os.Args[2]

	beforeTime := time.Now()
	for i := 0; i < numReq; i++ {
		whatever := "whatever"
		pingStr := "ping"
		ping := &pbs.Ping{
			Id:      &whatever,
			Payload: &pingStr,
		}
		pingBytes, err := proto.Marshal(ping)
		if err != nil {
			panic(err.Error())
		}
		req, err := http.NewRequest("GET", targetUrl, bytes.NewBuffer(pingBytes))
		if err != nil {
			panic(err.Error())
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err.Error())
		}
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		pong := &pbs.Pong{}
		if err := proto.Unmarshal(bs, pong); err != nil {
			if *pong.Payload != "pong" {
				panic("Expected pong")
			}
		}
		resp.Body.Close()
	}

	dur := time.Now().Sub(beforeTime)
	avgLatency := time.Duration(dur.Nanoseconds() / int64(numReq))
	fmt.Printf("Elapsed time: %s\n", dur)
	fmt.Printf("Mean latency: %s\n", avgLatency.String())
	fmt.Printf("Throughput: %.1f req/sec\n", float64(numReq)/dur.Seconds())
}
