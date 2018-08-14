package todos

import (
	"context"
	"io"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mhamrah/todos/gen"
	"github.com/oklog/ulid"
)

type IDGenerator func() string

func ulidGenerator(entropy io.Reader) IDGenerator {
	return func() string {
		return ulid.MustNew(ulid.Now(), entropy).String()
	}
}

type Server struct {
	storage Storage
	genID   IDGenerator
}

func NewServer(storage Storage) *Server {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Server{
		storage: storage,
		genID:   ulidGenerator(entropy),
	}
}

func (s *Server) ListTodos(ctx context.Context, in *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	result, err := s.storage.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListTodosResponse{Todos: result}, nil
}

func (s *Server) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.Todo, error) {
	return s.storage.Read(ctx, in.Id)
}

func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.Todo, error) {
	input := in.Todo
	input.Id = s.genID()

	err := s.storage.Save(ctx, input)
	if err != nil {
		return nil, err
	}

	return input, nil

}

func (s *Server) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*empty.Empty, error) {
	return nil, s.storage.Delete(ctx, in.Id)
}

func (s *Server) DeleteAllTodos(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	return nil, s.storage.DeleteAll(ctx)
}

func (s *Server) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.Todo, error) {
	err := s.storage.Save(ctx, in.Todo)
	if err != nil {
		return nil, err
	}
	return in.Todo, nil
}
