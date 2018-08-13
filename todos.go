package todos

import (
	"context"

	pb "github.com/mhamrah/todos/gen"
)

type Server struct {
}

func (s *Server) ListTodos(ctx context.Context, in *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	return &pb.ListTodosResponse{
		Todos: []*pb.Todo{
			&pb.Todo{Title: "foo"},
		},
	}, nil
}
