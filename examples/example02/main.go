package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	UnimplementedGreeterServer
}

func (s service) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("received %v", in.GetName())

	return &HelloReply{
		Message: fmt.Sprintf("Hello %s", in.GetName()),
	}, nil
}

func main() {
	server := grpc.NewServer()

	RegisterGreeterServer(server, service{})

	reflection.Register(server)

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	log.Printf("listening at %v", lis.Addr())

	server.Serve(lis)
}
