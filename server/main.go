package main

import (
	"net"
	"google.golang.org/grpc/grpclog"
	pb "./proto"
	"google.golang.org/grpc"
	"log"
	"fmt"
	"context"
	"io/ioutil"
)

const (
	Address = ":8765"
)

func main() {
	log.Println("=====================")
	log.Println("Starting server on port" + Address)
	log.Println("=====================")

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUploadFileServiceServer(grpcServer, newServer())

	grpcServer.Serve(listen)
}

type uploadFileServiceServer struct {
}

func (s *uploadFileServiceServer) UploadFile(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println(fmt.Sprintf("Username = %s", req.UserName))

	err := ioutil.WriteFile("file.jpg", req.File, 0644); if err != nil {
		fmt.Println(err.Error())
	}

	return &pb.Response{IsSuccessful: true}, nil
}

func newServer() *uploadFileServiceServer {
	return &uploadFileServiceServer{}
}