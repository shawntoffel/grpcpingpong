package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/shawntoffel/grpcpingpong/pingpong"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewPingPongClient(conn)

	reply, err := client.Ping(context.Background(), &pb.PingRequest{})

	if err != nil {
		log.Fatalf("Could not ping: %v", err)
	}

	log.Printf(reply.Pong)
}
