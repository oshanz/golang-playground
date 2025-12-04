package main

import (
	"log"
	"net"

	pb "github.com/oshanz/grpc-hello/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var address string = "0.0.0.0:50051"

func main() {
	server, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	log.Printf("listening..")

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(server); err != nil {
		log.Fatalf("GRPC failed %v", err)
	}
}
