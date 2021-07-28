package main

import (
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chatpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %s/n", err)
	}
	defer conn.Close()
	c := chatpb.NewChatServiceClient(conn)
	message := chatpb.Message{
		Body: "hello from client",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when callin sayhello %s", err)
	}
	log.Printf("response from server\n %s", response.Body)
}
