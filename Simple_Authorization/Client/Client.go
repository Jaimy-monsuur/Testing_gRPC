package main

import (
	"Simple_gRPC/Simple_Authorization/Proto/Authpd"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

// Define your API key here.
const apiKey = "43543634098634--60346"

func main() {

	/*
		// Load the server's CA certificate (you can also use Insecure if it's a self-signed certificate).
		creds, err := credentials.NewClientTLSFromFile("server.crt", "")
		if err != nil {
		log.Fatalf("failed to load server CA certificate: %v", err)
		}

		// Create a gRPC connection to the server using transport credentials.
		conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	*/

	// Create a gRPC connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := Authpd.NewGreetServiceClient(conn)

	// Create a context with a deadline for the RPC call.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a request message.
	req := Authpd.GreetRequest{
		FirstName: "Jaimy",
		LastName:  "Monsuur",
	}

	// Add the API key to the request metadata.
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", apiKey))

	// Call the Greet RPC.
	res, err := client.Greet(ctx, &req)
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	// Print the response.
	fmt.Println(res.GetMessage())
}
