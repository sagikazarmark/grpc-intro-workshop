package main

import (
	"fmt"

	example01 "github.com/sagikazarmark/grpc-intro-workshop/solutions/example01/generated/buf/go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	person := &example01.Person{
		Name:  "John Doe",
		Id:    1,
		Email: "john.doe@example.com",
	}

	fmt.Printf("String: %s\n", person)

	b, err := proto.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wire format: %v\n", b)

	b, err = protojson.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON: %s\n", b)
}
