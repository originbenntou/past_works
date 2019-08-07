package main

import (
	"google.golang.org/grpc"
	pb "grpc_go_sample/echo/proto"
	"log"
	"net"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()
	pb.RegisterEchoServiceServer(srv, &echoService{})

	log.Printf("start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
