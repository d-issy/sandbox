package main

import (
	"context"
	"log"
	"time"

	pb "github.com/d-issy/sandbox/grpc/calc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func add(c pb.CalcServiceClient, ctx context.Context, x int32, y int32) {
	res, err := c.Add(ctx, &pb.AddRequest{First: x, Second: y})
	if err != nil {
		log.Fatalf("cloud not add: %v\n", err)
	}
	log.Printf("received result: %v", res.Result)
}

func sub(c pb.CalcServiceClient, ctx context.Context, x int32, y int32) {
	res, err := c.Sub(ctx, &pb.SubRequest{First: x, Second: y})
	if err != nil {
		log.Fatalf("cloud not add: %v\n", err)
	}
	log.Printf("received result: %v", res.Result)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	c := pb.NewCalcServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	add(c, ctx, 1, 2)
	add(c, ctx, 30, 40)
	sub(c, ctx, 5, 2)
	sub(c, ctx, 25, 35)

}
