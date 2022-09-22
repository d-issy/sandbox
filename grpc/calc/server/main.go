package main

import (
	context "context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/d-issy/sandbox/grpc/calc/protobuf"
)

var addr string = "localhost:50051"

type server struct {
	pb.UnimplementedCalcServiceServer
}

func (*server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("invoked add(%d, %d)\n", in.First, in.Second)
	return &pb.AddResponse{Result: in.First + in.Second}, nil
}

func (*server) Sub(ctx context.Context, in *pb.SubRequest) (*pb.SubResponse, error) {
	log.Printf("invoked sub(%d, %d)\n", in.First, in.Second)
	return &pb.SubResponse{Result: in.First - in.Second}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listen to %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
