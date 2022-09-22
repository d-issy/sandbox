package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/d-issy/sandbox/grpc/tips/protobuf"
)

var addr string = "localhost:50051"

type server struct {
	pb.UnimplementedTipsServiceServer
}

func (*server) DoError(context.Context, *pb.Request) (*pb.Response, error) {
	log.Printf("DoError invoked")
	return nil, status.Error(codes.PermissionDenied, "permission error")
}

func (*server) DoWithDeadline(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("DoWithDeadline invoked")
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("The client canceld the request!")
		return nil, status.Error(codes.Canceled, "the client canceled")
	} else {
		log.Println("The server canceld the request!")
		return nil, status.Error(codes.DeadlineExceeded, "the server canceled")
	}
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listen to %v\n", addr)
	s := grpc.NewServer()
	pb.RegisterTipsServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
