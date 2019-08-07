package main

import (
	pb "grpc_go_sample/upload/proto"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
)

type fileService struct{}

func (s *fileService) Upload(stream pb.FileService_UploadServer) error {
	var name string
	var blob []byte
	for {
		c, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("complete %d bytes\n", len(blob))
				break
			}
			panic(err)
		}
		name = c.GetName()
		blob = append(blob, c.GetData()...)
	}
	//name = req.GetName()
	fp := filepath.Join("../resource", name)
	ioutil.WriteFile(fp, blob, 0644)

	log.Printf("complete %v\n", fp)

	stream.SendAndClose(&pb.FileResponse{Size: int64(len(blob))})

	return nil
}
