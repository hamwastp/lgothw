package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 跨语言的rpc通信
type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
