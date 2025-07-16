package client_call

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}
	fmt.Println("Response", res)
}
