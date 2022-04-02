package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Head("http://localhost:50510")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status: ", res.Status)
	log.Println("StatusCode: ", res.StatusCode)
	log.Println("Headers: ", res.Header)
	log.Println("Content-Length: ", res.Header.Get("Content-Length"))
}
