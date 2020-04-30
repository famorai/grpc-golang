package main

import (
	"context"
	"log"

	"github.com/famorai/grpc-hello-go/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "Fabio Morais",
	}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during the execution: %v", err)
	}

	log.Println(res.Msg)
}
