package service

import (
	"context"
	"fmt"

	"github.com/ersa97/latihan-grpc/server/chat"
)

type SayService struct {
}

func (s *SayService) SayHello(context context.Context, request *chat.SayRequest) (*chat.SayResponse, error) {
	response := fmt.Sprintf("Message Received, message is %s", request.Word)
	return &chat.SayResponse{
		Content: response,
	}, nil
}
