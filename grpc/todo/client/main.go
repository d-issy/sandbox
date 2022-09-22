package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/d-issy/sandbox/grpc/todo/protobuf"
)

var (
	addr = flag.String("addr", "localhost:50051", "server addr")
)

func input(prompt string) string {
	fmt.Printf("%s: ", prompt)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func get(client pb.TodoServiceClient) (*pb.Todo, error) {
	id := input("id")
	todo, err := client.Get(context.Background(), &pb.TodoId{Id: id})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func show(todo *pb.Todo) {
	log.Printf("TODO: id: \"%s\" title: \"%s\", completed: %v\n", todo.Id, todo.Title, todo.Completed)
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect to todo service: %v\n", err)
	}

	client := pb.NewTodoServiceClient(conn)

	cmd := input("command")
	switch cmd {
	case "list":
		res, err := client.List(context.Background(), &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not execute list command: %v\n ", err)
		}
		for {
			todo, err := res.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error from receiving todo: %v\n", err)
			}
			show(todo)
		}
	case "get":
		todo, err := get(client)
		if err != nil {
			log.Fatalln("could not create todo")
		}
		show(todo)
	case "create":
		title := input("title")
		_, err := client.Create(context.Background(), &pb.Todo{Title: title})
		if err != nil {
			log.Fatalln("could not create todo")
		}
	case "update":
		todo, err := get(client)
		if err != nil {
			log.Fatalf("not found todo %v\n", err)
		}
		title := input("title")
		if title == "" {
			title = todo.Title
		}
		completed, err := strconv.ParseBool(input("completed"))
		if err != nil {
			completed = todo.Completed
		}
		_, err = client.Update(context.Background(), &pb.Todo{Id: todo.Id, Title: title, Completed: completed})
		if err != nil {
			log.Fatalln("could not update todo")
		}
	case "delete":
		todo, err := get(client)
		if err != nil {
			log.Fatalf("not found todo %v\n", err)
		}
		_, err = client.Delete(context.Background(), &pb.TodoId{Id: todo.Id})
		if err != nil {
			log.Fatalf("could not delete todo: %v\n", todo)
		}
	default:
		log.Fatalf("not found command: %v\n", cmd)
	}
}
