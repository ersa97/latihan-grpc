package service

import (
	"context"
	"fmt"
	"log"

	"github.com/ersa97/latihan-grpc/server/chat"
)

type SayService struct {
}

func (s *SayService) SayHello(context context.Context, request *chat.SayRequest) (*chat.SayResponse, error) {
	response := fmt.Sprintf("Message Received, message is %s", request.Word)
	log.Printf("Get SayRequest : %s\n", request.GetWord())
	return &chat.SayResponse{
		Content: response,
	}, nil
}
