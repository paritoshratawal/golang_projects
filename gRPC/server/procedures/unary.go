package procedures

import (
	"context"

	pb "github.com/paritoshratawal/golang_projects/gRPC/proto"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		message: "Hello",
	}, nil
}
