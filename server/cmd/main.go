package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mhamrah/todos/gen"
	"github.com/mhamrah/todos/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTodosServer(grpcServer, todos.NewServer(todos.MemoryStorage{}))

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
