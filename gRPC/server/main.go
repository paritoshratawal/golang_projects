package main

import (
	"context"
	"log"
	"net"

	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type HelloServer struct {
	pb.GreetServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start tcp server: %v", err)
	}
	grpc_server := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpc_server, &HelloServer{})
	log.Printf("server started at :%v", listner.Addr())
	if err = grpc_server.Serve(listner); err != nil {
		log.Fatalf("Failed to start grpc: %v", err)
	}
}
