package main

// 4.4.4 发布和订阅模式

import (
	s "go.learning.com/ch4/protobuf/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	s.RegisterPubsubServiceServer(grpcServer, s.NewPubsubService()) // 注意不能直接new, 否则成员会出现空指针

	lis, err := net.Listen("tcp", ":1236")
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
