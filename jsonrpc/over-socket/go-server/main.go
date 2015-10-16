package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type RPCFunc uint8

func (*RPCFunc) Echo(arg *string, result *string) error {
	log.Print("Arg passed: " + *arg)
	*result = ">" + *arg + "<"
	return nil
}

func main() {
	log.Println("Starting RPC Server...")
	l, err := net.Listen("tcp", "localhost:9394")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer l.Close()
	log.Print("Listening on: ", l.Addr())
	rpc.Register(new(RPCFunc))
	for {
		log.Println("Waiting for connections ...")
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Accept error: %s\n", err)
			continue
		}
		log.Printf("Connection started: %v\n", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
