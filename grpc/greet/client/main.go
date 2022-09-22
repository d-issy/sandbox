package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/d-issy/sandbox/grpc/greet/protobuf"
)

var addr string = "localhost:50051"

func greet(c pb.GreetServiceClient, ctx context.Context, name string) {
	res, err := c.Greet(ctx, &pb.GreetRequest{Name: name})
	if err != nil {
		log.Fatalf("Cloud not greatMany: %v\n", err)
	}
	log.Printf("received from server: %s\n", res.Result)
}

func greetMany(c pb.GreetServiceClient, ctx context.Context, name string) {
	stream, err := c.GreetManyTimes(ctx, &pb.GreetRequest{Name: name})
	if err != nil {
		log.Fatalf("Cloud not GreetMany: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while reading the stream: %v\n", err)
		}
		log.Printf("received from server: %s\n", res.Result)
	}
	stream.CloseSend()
}

func longGreet(c pb.GreetServiceClient, ctx context.Context, names []string) {
	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("Cloud not LongGreet: %v\n", err)
	}
	for _, name := range names {
		stream.Send(&pb.GreetRequest{Name: name})
		time.Sleep(200 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}
	log.Printf("LongGreet: %s\n", res.Result)
}

func greetEveryone(c pb.GreetServiceClient, ctx context.Context, names []string) {
	stream, err := c.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("Cloud not GreetEveryone: %v\n", err)
	}

	waitc := make(chan struct{})

	// sending
	go func() {
		for _, name := range names {
			log.Printf("greet request: %s", name)
			if err := stream.Send(&pb.GreetRequest{Name: name}); err != nil {
				log.Printf("Error while sending: %s\n", err)
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// reciving
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("received: %v", in.Result)
		}
		waitc <- struct{}{}
	}()

	<-waitc
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		greet(c, ctx, "world")
		wg.Done()
	}()

	go func() {
		greetMany(c, ctx, "world")
		wg.Done()
	}()

	go func() {
		longGreet(c, ctx, []string{"Taro", "Hanako"})
		wg.Done()
	}()

	go func() {
		greetEveryone(c, ctx, []string{"Taro", "Hanako"})
		wg.Done()
	}()

	wg.Wait()
}
