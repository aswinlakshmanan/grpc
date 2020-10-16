package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tutorial/go-grpc-tutorial/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)

	}
	defer conn.Close()

	c := chat.NewchatServiceClient(conn)
	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHello %s", err)

	}

	log.Printf("Response from the server : %s", response.Body)
}
