package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/d-issy/sandbox/grpc/greet/protobuf"
)

var addr string = "localhost:50051"

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greeat invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.Name}, nil
}

func (*server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTyimes invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Hello %s (%d)", in.Name, i)
		stream.Send(&pb.GreetResponse{Result: message})
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet invoked")
	msg := "Hello"
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading client: %v\n", err)
		}
		log.Printf("with %s\n", in.Name)
		msg += fmt.Sprintf(", %s", in.Name)
	}
	return stream.SendAndClose(&pb.GreetResponse{Result: msg + "!!"})
}

func (*server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone invoked")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while reading client stream %v\n", err)
			break
		}
		log.Printf("received and send: %s\n", in.Name)
		err = stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello %s", in.Name)})
		if err != nil {
			log.Printf("Error while sending client stream %v\n", err)
			break
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listen to %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
