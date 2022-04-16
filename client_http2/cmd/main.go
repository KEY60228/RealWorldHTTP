package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("Protocol Version: %s\n", res.Proto)
}