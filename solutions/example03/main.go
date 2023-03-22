package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := NewGreeterClient(conn)

	name := "HWSW"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	req := &HelloRequest{
		Name: name,
	}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", resp.GetMessage())
}
