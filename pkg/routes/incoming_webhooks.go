package routes

import (
	"context"
	"fmt"

	iw_pb "snack/pkg/grpc/incoming_webhooks"
)

type incomingWebhooksServer struct {
	iw_pb.UnimplementedIncomingWebhooksServiceServer
}

func (server *incomingWebhooksServer) Register(ctx context.Context, req *iw_pb.RegisterRequest) (*iw_pb.RegisterResponse, error) {
	return &iw_pb.RegisterResponse{
		Message: fmt.Sprintf("response"),
	}, nil
}

func NewIncomingWebhooksServer() *incomingWebhooksServer {
	return &incomingWebhooksServer{}
}
