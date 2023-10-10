package main

import (
	"context"
	"log"
	"net"

	"Simple_gRPC/greet/proto/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

// Greet greets with FirstName
func (*server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	result := "Hello " + in.GetGreeting().GetFirstName() + " " + in.GetGreeting().GetLastName()
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}
