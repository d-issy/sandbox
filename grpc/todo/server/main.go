package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/d-issy/sandbox/grpc/todo/protobuf"
)

var (
	addr = flag.String("addr", "localhost:50051", "server addr")
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer

	todos map[string]Todo
}

func (s *TodoService) List(_ *emptypb.Empty, stream pb.TodoService_ListServer) error {
	for _, t := range s.todos {
		if err := stream.Send(&pb.Todo{
			Id:        t.Id,
			Title:     t.Title,
			Completed: t.Completed,
		}); err != nil {
			log.Fatalln("can not send todo data")
		}
	}
	return nil
}

func (s *TodoService) Get(ctx context.Context, in *pb.TodoId) (*pb.Todo, error) {
	t, ok := s.todos[in.Id]
	if ok {
		return &pb.Todo{
			Id:        t.Id,
			Title:     t.Title,
			Completed: t.Completed,
		}, nil
	}
	return nil, status.Error(codes.NotFound, "todo not found")
}

func (s *TodoService) Create(ctx context.Context, in *pb.Todo) (*pb.TodoId, error) {
	log.Printf("create todo: %v\n", in)
	t := Todo{
		Id:        uuid.New().String(),
		Title:     in.Title,
		Completed: false,
	}
	s.todos[t.Id] = t
	return &pb.TodoId{Id: t.Id}, nil
}

func (s *TodoService) Update(ctx context.Context, in *pb.Todo) (*emptypb.Empty, error) {
	if _, ok := s.todos[in.Id]; ok {
		s.todos[in.Id] = Todo{
			Id:        in.Id,
			Title:     in.Title,
			Completed: in.Completed,
		}
		return &emptypb.Empty{}, nil
	}
	return nil, status.Error(codes.NotFound, "todo not found")
}

func (s *TodoService) Delete(ctx context.Context, in *pb.TodoId) (*emptypb.Empty, error) {
	if _, ok := s.todos[in.Id]; ok {
		delete(s.todos, in.Id)
		return &emptypb.Empty{}, nil
	}
	return nil, status.Error(codes.NotFound, "todo not found")
}

func newServer() *TodoService {
	return &TodoService{todos: make(map[string]Todo)}
}

type Todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to Start Listening: %v", err)
	}
	log.Printf("Listening to %s...", *addr)

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, newServer())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
