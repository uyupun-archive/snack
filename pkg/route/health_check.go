package routes

import (
	"context"
	"fmt"

	hc_pb "snack/pkg/grpc/health_check"
)

type healthCheckServer struct {
	hc_pb.UnimplementedHealthCheckServiceServer
}

func (server *healthCheckServer) Echo(ctx context.Context, req *hc_pb.PingRequest) (*hc_pb.PongResponse, error) {
	return &hc_pb.PongResponse{
		Message: fmt.Sprintf("pong, %s!", req.GetName()),
	}, nil
}

func NewHealthCheckServer() *healthCheckServer {
	return &healthCheckServer{}
}
