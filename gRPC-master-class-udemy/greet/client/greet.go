package main

import (
	"context"
	"log"

	pb "github.com/oshanz/grpc-hello/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Oshan",
	})
	if err != nil {
		log.Fatalf("%a", err)
	}
	log.Printf("response %s", res.Result)
}
