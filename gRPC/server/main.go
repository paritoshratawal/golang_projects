package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

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

func (s *HelloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello! " + name,
		}
		if err := stream.Send(res); err != nil {
			// log.Fatalf("Error while sending stream data to client")
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *HelloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}

	return nil
}

func (s *HelloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	// var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
			// return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello! " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while sending stream data to client")
			return err
		}
		// time.Sleep(1 * time.Second)

		// messages = append(messages, "Hello", req.Name)
	}

	return nil
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
