package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "greeter-grpc-byoc/proto"
)

type GreetResponse struct {
	Message string `json:"message"`
}

func greetHandler(client pb.GreetingServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "name parameter is required", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Printf("gRPC call failed: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := GreetResponse{Message: resp.GetMessage()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	target := os.Getenv("GRPC_SERVER_TARGET")
	if target == "" {
		target = "localhost:50051"
	}

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)

	http.HandleFunc("/greet", greetHandler(client))

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting HTTP server on port %s", port)
	log.Printf("gRPC backend: %s", target)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
