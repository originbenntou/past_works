package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc_go/helloworld"
)

type service struct{}

func (s *service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(ctx, req)
	log.Println("call from", req.Yamada)
	rsp := new(pb.HelloReply)
	rsp.Message = "Hello " + req.Yamada + "."
	return rsp, nil
}

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service{})
	s.Serve(l)
}
