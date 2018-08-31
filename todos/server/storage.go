package todos

import (
	"context"
	"errors"

	pb "github.com/mhamrah/grpc-example/gen"
)

type Storage interface {
	Read(context.Context, string) (*pb.Todo, error)
	ReadAll(context.Context) ([]*pb.Todo, error)
	Delete(context.Context, string) error
	DeleteAll(context.Context) error
	Save(context.Context, *pb.Todo) error
}

type MemoryStorage map[string]*pb.Todo

func (m MemoryStorage) Read(ctx context.Context, id string) (*pb.Todo, error) {
	r, exists := m[id]
	if !exists {
		return nil, errors.New("not found")
	}
	return r, nil
}

func (m MemoryStorage) ReadAll(ctx context.Context) ([]*pb.Todo, error) {
	var r []*pb.Todo

	for _, t := range m {
		r = append(r, t)
	}

	return r, nil
}

func (m MemoryStorage) Delete(ctx context.Context, id string) error {
	delete(m, id)
	return nil
}

func (m MemoryStorage) DeleteAll(ctx context.Context) error {

	for k := range m {
		delete(m, k)
	}

	return nil
}

func (m MemoryStorage) Save(ctx context.Context, todo *pb.Todo) error {
	m[todo.Id] = todo
	return nil
}
