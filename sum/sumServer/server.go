package main

import (
	"Simple_gRPC/sum/proto/sumpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	sumpb.UnimplementedSumServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to liesten: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}

// Add returns sum of two integers
func (*server) Add(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	a, b := req.GetNumbers().GetA(), req.GetNumbers().GetB()
	sum := a + b
	return &sumpb.SumResponse{Result: sum}, nil
}
