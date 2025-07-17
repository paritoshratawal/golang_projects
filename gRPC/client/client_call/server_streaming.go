package client_call

import (
	"context"
	"io"
	"log"

	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
)

func CallSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Starts recieving server stream")
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("Error while streaming from server: %v", err)
	}

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
	log.Printf("Streaming finished")
}
