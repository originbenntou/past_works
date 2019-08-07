package main

import (
	"context"
	pb "grpc_go_sample/echo/proto"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponce, error) {
	return &pb.EchoResponce{
		Message: req.GetMessage(),
	}, nil
}
