package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	UnimplementedGreeterServer
}

func (s service) SayHello(in *HelloRequest, stream Greeter_SayHelloServer) error {
	log.Printf("received: %v", in.GetName())

	for i := 0; i < 5; i++ {
		resp := &HelloReply{Message: fmt.Sprintf("Hello %s", in.GetName())}
		if i > 0 {
			resp = &HelloReply{Message: fmt.Sprintf("Hello again %s", in.GetName())}
		}

		stream.Send(resp)

		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	server := grpc.NewServer()

	// Register Greeter service implementation on the gRPC server.
	RegisterGreeterServer(server, service{})

	// Register reflection service on the gRPC server.
	reflection.Register(server)

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
