// A painfully simple latency benchmark for a simple protobuf-over-HTTP service.
// This is intended to show the best case latency for such a service.
package main

import (
	pbs ".."
	"code.google.com/p/goprotobuf/proto"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		ping := &pbs.Ping{}
		if err := proto.Unmarshal(bs, ping); err != nil {
			panic(err.Error())
		}
		if *ping.Payload != "ping" {
			panic("Expected payload \"ping\"")
		}
		r.Body.Close()
		w.Header().Set("Content-type", "text/plain")

		whatever := "whatever"
		pongStr := "pong"
		pong := &pbs.Pong{
			Id:      &whatever,
			Payload: &pongStr,
		}
		pongBytes, err := proto.Marshal(pong)
		if err != nil {
			panic(err.Error())
		}
		if _, err := w.Write(pongBytes); err != nil {
			panic(err.Error())
		}
	})

	err := http.ListenAndServe("0.0.0.0:4321", nil)
	if err != nil {
		panic(err.Error())
	}
}
