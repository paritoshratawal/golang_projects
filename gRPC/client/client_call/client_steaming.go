package client_call

import (
	"context"
	"log"

	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
)

func CallSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client started streaming")
	stream, err := client.SayHelloClientStreaming(context.TODO())
	if err != nil {
		log.Fatalf("Error while streaming to server: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending stream data to server: %v", err)
		}
		log.Println("Send the request with name: ", name)
	}

}
