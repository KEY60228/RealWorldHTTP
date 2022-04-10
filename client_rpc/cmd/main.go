package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:50510")
	if err != nil {
		panic(err)
	}

	var result int
	a := 4
	b := 5
	args := &Args{a, b}
	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("%d * %d = %d\n", a, b, result)
}
