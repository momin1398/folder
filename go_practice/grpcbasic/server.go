package main

import (
	"log"
	"net"

	"github.com/tutorialedge/go-grpc-tutorial/chatpb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}
	s := chatpb.Server{}

	grpcServer := grpc.NewServer()
	chatpb.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc port 8080: %v", err)

	}

}
