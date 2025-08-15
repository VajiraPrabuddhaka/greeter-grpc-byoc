package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "greeter-grpc-byoc/proto"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (s *server) SayHelloStream(req *pb.HelloRequest, stream pb.GreetingService_SayHelloStreamServer) error {
	for i := 0; i < 5; i++ {
		response := &pb.HelloResponse{
			Message: fmt.Sprintf("Hello #%d, %s!", i+1, req.GetName()),
		}
		if err := stream.Send(response); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
