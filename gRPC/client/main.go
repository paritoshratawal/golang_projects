package main

import (
	"log"

	pc "github.com/paritoshratawal/golang_projects/gRPC/client/client_call"
	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "8080"
)

func main() {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"Ashu", "Paritosh", "Rashima", "Sheetal"},
	// }
	pc.CallSayHello(client)
}
