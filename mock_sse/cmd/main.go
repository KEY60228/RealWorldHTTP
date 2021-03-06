package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

var html []byte

func handlerHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write(html)
}

func handlerPrimeSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var num int64 = 1
	for id := 1; id <= 100; id++ {
		select {
		case <-ctx.Done():
			fmt.Println("Connection closed from client")
			return
		default:
			// do nothing
		}
		for {
			num++
			if big.NewInt(num).ProbablyPrime(20) {
				fmt.Println(num)
				fmt.Fprintf(w, "data: {\"id\": %d, \"number\": %d}\n\n", id, num)
				flusher.Flush()
				time.Sleep(time.Second)
				break
			}
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Connection closed from server")
}

func main() {
	var err error
	html, err = ioutil.ReadFile("template/index.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/prime", handlerPrimeSSE)
	fmt.Println("start http listening :50510")
	err = http.ListenAndServe(":50510", nil)
	fmt.Println(err)
}
