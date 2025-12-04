package main

import (
	"context"
	"log"

	pb "github.com/oshanz/grpc-hello/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet %s \n", in)
	return &pb.GreetResponse{
		Result: "Hi " + in.FirstName,
	}, nil
}
