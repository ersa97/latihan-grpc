package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ersa97/latihan-grpc/server/chat"
	"github.com/ersa97/latihan-grpc/server/service"
	"google.golang.org/grpc"
)

func main() {
	//create network
	conn, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		log.Fatal("Can't Create Network")
	}
	//create grpc
	server := grpc.NewServer()
	//register implementation into GRPC server
	chat.RegisterHelloServiceServer(server, &service.SayService{})

	fmt.Println("GRPC server is running at 0.0.0.0:8000")

	if err := server.Serve(conn); err != nil {
		log.Fatal("Can't Create GRPC Server")
	}
}
