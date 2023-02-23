package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	hc_pb "snack/pkg/grpc/health_check"
)

func main() {
	fmt.Println("start gRPC client")

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed")
		return
	}
	defer conn.Close()

	client := hc_pb.NewHealthCheckServiceClient(conn)
	name := "snack"
	req := &hc_pb.PingRequest{
		Name: name,
	}
	res, err := client.Echo(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}
