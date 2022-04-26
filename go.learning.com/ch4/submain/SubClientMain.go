package main

// 4.4.4 发布和订阅模式
import (
	"context"
	"fmt"
	"io"
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
	stream, err := client.Subscribe(
		context.Background(), &v.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}
