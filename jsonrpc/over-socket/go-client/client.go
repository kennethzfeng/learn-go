package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9394")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	client := jsonrpc.NewClient(conn)
	var reply string
	for i, arg := range os.Args {
		if i == 0 {
			continue
		} // ignore program name
		fmt.Printf("Sending: %s\n", arg)
		client.Call("RPCFunc.Echo", arg, &reply)
		fmt.Printf("Reply: %s\n", reply)
	}
}
