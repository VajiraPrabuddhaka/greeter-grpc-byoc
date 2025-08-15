package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "greeter-grpc-byoc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	stream, err := c.SayHelloStream(context.Background(), &pb.HelloRequest{Name: "Stream User"})
	if err != nil {
		log.Fatalf("could not get stream: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive: %v", err)
		}
		log.Printf("Stream Greeting: %s", response.GetMessage())
	}
}