package proto

import (
	context "context"
	fmt "fmt"
	"io"
	"log"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// 跨语言的rpc通信

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	log.Println("Hello...")
	reply.Value = "Hello:" + request.GetValue()
	return nil
}

type Authentication struct {
	Login    string
	Password string
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	if val, ok := md["login"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != "gopher" || appkey != "password" {
		return grpc.Errorf(codes.Unauthenticated, "invalid token: appid=%s, appkey=%s", appid, appkey)
	}

	return nil
}

type HelloServiceImpl struct {
	auth Authentication
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {

	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
