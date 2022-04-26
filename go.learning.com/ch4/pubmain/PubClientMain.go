package main

import (
	"context"
	"log"

	// docker项目中提供了一个pubsub的极简实现

	v "go.learning.com/ch4/protobuf/pubsub"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1236", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &v.String{Value: "golang: hello Go"})

	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &v.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
