package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shawntoffel/grpcpingpong/pingpong"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Ping(ctx context.Context, pingRequest *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Pong: "pong"}, nil
}

func main() {
	s := grpc.NewServer([]grpc.ServerOption{}...)

	pb.RegisterPingPongServer(s, &server{})

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
