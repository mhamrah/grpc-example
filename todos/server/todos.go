package todos

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/mhamrah/grpc-example/gen"
	"github.com/oklog/ulid"
	"go.uber.org/zap"
)

type IDGenerator func() string

func ulidGenerator(entropy io.Reader) IDGenerator {
	return func() string {
		return ulid.MustNew(ulid.Now(), entropy).String()
	}
}

type Server struct {
	pb.UnimplementedTodosServer
	storage Storage
	genID   IDGenerator
	logger  *zap.SugaredLogger
}

func NewServer(storage Storage, logger *zap.Logger) *Server {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Server{
		storage: storage,
		genID:   ulidGenerator(entropy),
		logger:  logger.Sugar(),
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
	s.logger.With("id", in.Id).Info("Get Todo")
	return s.storage.Read(ctx, in.Id)
}

func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.Todo, error) {
	input := in.Todo
	input.Id = s.genID()

	f, err := os.Open(os.DevNull)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		fmt.Fprintf(f, ".")
	}

	err = s.storage.Save(ctx, input)
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
