package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc_go_sample/upload/proto"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &fileService{})
	if err = s.Serve(lis); err != nil {
		log.Printf("failed to listen: %v\n", err)
	}
}
