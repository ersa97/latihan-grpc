package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ersa97/latihan-grpc/server/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to connect GRPC server")
	}
	defer conn.Close()

	server := chat.NewHelloServiceClient(conn)

	response, err := server.SayHello(context.Background(), &chat.SayRequest{Word: "Hello Server"})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response :\n%s\n", response.GetContent())

}
