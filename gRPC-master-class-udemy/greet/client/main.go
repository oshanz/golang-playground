package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/oshanz/grpc-hello/greet/proto"
)

var address string = "0.0.0.0:50051"

func main() {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer connection.Close()

	c := pb.NewGreetServiceClient(connection)
	doGreet(c)
}
