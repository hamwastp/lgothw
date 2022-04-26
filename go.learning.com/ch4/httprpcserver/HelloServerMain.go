package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	hello "go.learning.com/ch4/protobuf/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func filter1(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("fileter:", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return handler(ctx, req)
}

func filter2(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("fileter:", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return handler(ctx, req)
}

func main() {
	// 指定使用服务端证书创建一个 TLS credentials
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	// 指定使用 TLS credentials
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(filter1, filter2)))
	hello.RegisterHelloServiceServer(grpcServer, new(hello.HelloServiceImpl))

	// lis, err := net.Listen("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 用于gocurl工具提供注册信息查询服务
	reflection.Register(grpcServer)

	// grpcServer.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello world")
		//log.Fatal("11111")
	})

	http.ListenAndServeTLS(":1234", "server.crt", "server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 如果将gRPC和Web服务放在⼀起，会导致gRPC和Web路径的冲突，在处理时我们需要区分两类服务。
			// 每个gRPC调⽤请求的Content-Type类型会被标注为"application/grpc"类型
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}),
	)
}
