package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", "http://localhost:50510", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}
