package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	hc_pb "snack/pkg/grpc/health_check"
	iw_pb "snack/pkg/grpc/incoming_webhooks"
	routes "snack/pkg/routes"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	hc_pb.RegisterHealthCheckServiceServer(server, routes.NewHealthCheckServer())
	iw_pb.RegisterIncomingWebhooksServiceServer(server, routes.NewIncomingWebhooksServer())

	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server (port: %v)", port)
		server.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server ...")
	server.GracefulStop()
}
