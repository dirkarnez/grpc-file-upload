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
	"encoding/json"
	 b64 "encoding/base64"
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

type jsonOutput struct {
	UserName string
	File string
}

func (s *uploadFileServiceServer) UploadFile(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println(fmt.Sprintf("Username = %s", req.UserName))

	err := ioutil.WriteFile("file.jpg", req.File, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

 	j := &jsonOutput{UserName: req.UserName, File: b64.StdEncoding.EncodeToString(req.File)}
	b, err := json.Marshal(j)
    if err != nil {
		fmt.Println(err.Error())
		return nil, err
    }
    err = ioutil.WriteFile("output.json", b, 0644)
    if err != nil {
		fmt.Println(err.Error())
		return nil, err
    }
	return &pb.Response{IsSuccessful: true}, nil
}

func newServer() *uploadFileServiceServer {
	return &uploadFileServiceServer{}
}