package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

type Arith uint8

func (*Arith) Multiply(args *Args, reply *int) error {
	*reply = args.X * args.Y
	return nil
}

func startServer() {
	arith := new(Arith)

	server := rpc.NewServer()
	server.Register(arith)

	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func main() {
	go startServer()

	conn, err := net.Dial("tcp", "localhost:8222")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	args := &Args{7, 8}
	var reply int

	c := jsonrpc.NewClient(conn)

	for i := 0; i < 1; i++ {

		err = c.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.X, args.Y, reply)
	}
}
