package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:50510/chunked")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}
