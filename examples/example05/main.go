package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
)

type service struct {
	UnimplementedGreeterServer
}

func (s service) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	if token := md.Get("token"); len(token) != 1 || token[0] != "root" {
		// return nil, status.Error(codes.Unauthenticated, "unauthenticated")

		s := status.New(codes.Unauthenticated, "unauthenticated")

		err := &HelloError{
			Reason: "invalid token",
		}

		s, _ = s.WithDetails(err)

		return nil, s.Err()
	}

	log.Printf("received: %v", in.GetName())

	return &HelloReply{Message: fmt.Sprintf("Hello %s", in.GetName())}, nil
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
