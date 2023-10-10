package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"Simple_gRPC/Simple_Authorization/Proto/Authpd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Define your API key here.
const apiKey = "43543634098634--60346"

// GreetServiceServer implements the GreetService gRPC service.
type GreetServiceServer struct {
	Authpd.UnimplementedGreetServiceServer
}

func (s *GreetServiceServer) Greet(ctx context.Context, req *Authpd.GreetRequest) (*Authpd.GreetResponse, error) {
	// Authenticate the client using metadata (API key).
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	apiKeyHeader := md.Get("authorization")
	if len(apiKeyHeader) != 1 || apiKeyHeader[0] != apiKey {
		return nil, status.Error(codes.PermissionDenied, "invalid API key")
	}

	// If the client is authenticated, proceed with the greeting.
	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	message := fmt.Sprintf("Hello, %s %s!", firstName, lastName)

	return &Authpd.GreetResponse{Message: message}, nil
}

func main() {
	// Create a TCP listener on port 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server.
	server := grpc.NewServer()

	// Register the GreetServiceServer with the gRPC server.
	Authpd.RegisterGreetServiceServer(server, &GreetServiceServer{})

	log.Println("Starting gRPC server on :50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
