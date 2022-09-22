package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "github.com/d-issy/sandbox/grpc/tips/protobuf"
)

var addr string = "localhost:50051"

func doError(c pb.TipsServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.DoError(ctx, &pb.Request{})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Printf("grpc error: %v (code: %v)\n", s.Message(), s.Code())
		} else {

			log.Printf("other error: %v\n", err)
		}
	}
	log.Println("No Error")
}

func doWithDeadLine(c pb.TipsServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.DoWithDeadline(ctx, &pb.Request{})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Printf("grpc error: %v (code: %v)\n", s.Message(), s.Code())
		} else {

			log.Printf("other error: %v\n", err)
		}
	}
	log.Println("No Error")
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	c := pb.NewTipsServiceClient(conn)

	doError(c)
	doWithDeadLine(c)
}
