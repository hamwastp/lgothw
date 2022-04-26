package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	v "go.learning.com/ch4/protobuf/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// openssl genrsa -out server.key 2048
// openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" -key server.key -out server.crt

// 根证书: openssl genrsa -out ca.key 2048
// openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=gobook/CN=github.com" -key ca.key -out ca.crt

/*
# 生成.csr 证书签名请求文件
 openssl req -new -key server.key -out server.csr \
	-subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
	-reqexts SAN \
	-config <(cat /etc/pki/tls/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.server.grpc.io,DNS:*.refersmoon.com"))

# 签名生成.crt 证书文件
openssl x509 -req -days 3650 \
   -in server.csr -out server.crt \
   -CA client.crt -CAkey client.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /etc/pki/tls/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.server.grpc.io,DNS:*.refersmoon.com"))
*/

/*
使用根证书进行签名
openssl x509 -req -days 3650 \
   -in server.csr -out server.crt \
   -CA ca.crt -CAkey ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /etc/pki/tls/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.server.grpc.io,DNS:*.refersmoon.com"))
*/

// CA 证书
// 1) 生成.key  私钥文件
// genrsa 用于生成RSA私钥，不会生成公钥，因为公钥提取自私钥
// -out filename:将生成的私钥保存至filename文件,若未指定输出文件,则为标准输出。
// numbits:指定要生成的私钥的长度,默认为1024。该项必须为命令行的最后一项参数。
// openssl genrsa -out client.key 2048
//
// 生成.csr 证书签名请求文件
// 在申请证书之前，必须生成证书私钥和证书请求文件（Cerificate Signing Request，简称CSR）。
// CSR文件是公钥证书原始文件，包含了我们的服务器信息和单位信息，需要提交给CA认证中心进行审核。
// 2) openssl req -new -key client.key -out client.csr  -subj "/C=GB/L=China/O=grpc-client/CN=*.refersmoon.com"
// 自签名生成.crt 证书文件
// openssl req -new -x509 -days 3650 -key client.key -out client.crt  -subj "/C=GB/L=China/O=grpc-client/CN=*.refersmoon.com"

// GODEBUG=x509ignoreCN=0,gctrace=1  go run HelloClientMain.go

// 什么是根证书
// 为了避免证书的传递过程中被篡改，可以通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名。
// 这样客户端或者服务器在收到对方的证书后可以通过根证书进行验证证书的有效性

func main1() {
	// 客户端通过ca证书来验证服务的提供的证书
	creds, err := credentials.NewClientTLSFromFile(
		"client.crt", "hello.server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}

	// 建立连接时指定使用 TLS
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &v.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&v.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
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
	}()
}

type Authentication struct {
	Login    string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error) {
	return map[string]string{"Login": a.Login, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	// 客户端通过ca证书来验证服务的提供的证书
	certificate, err := tls.LoadX509KeyPair(
		"client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	// .cer/.crt是用于存放证书，它是2进制形式存放的，不含私钥。
	// .pem跟crt/cer的区别是它以Ascii来表示。
	// https://blog.csdn.net/qq_37049781/article/details/84837342
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		// ServerName:   tlsServerName, // Note: this is required!
		ServerName: "hello.server.grpc.io",
		RootCAs:    certPool,
	})

	auth := Authentication{
		Login:    "gopher",
		Password: "password",
	}

	// 建立连接时指定使用 TLS
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &v.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&v.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
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
	}()
}
