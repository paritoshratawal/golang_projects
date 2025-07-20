package client_call

import (
	"context"
	"io"
	"log"

	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
)

func recievingStream(waitCh chan struct{}, stream pb.GreetService_SayHelloBidirectionalStreamingClient) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Printf("Message: %v", msg)
	}
	close(waitCh)
}

func CallSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming to server streaming")
	stream, err := client.SayHelloBidirectionalStreaming(context.TODO())
	if err != nil {
		log.Fatalf("Error while streaming to server: %v", err)
	}

	waitCh := make(chan struct{})

	go recievingStream(waitCh, stream)

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending stream data to server: %v", err)
		}
		log.Println("Send the request with name: ", name)
	}
	stream.CloseSend()
	<-waitCh
	log.Println("Bidirectional Streaming Completed")
}
