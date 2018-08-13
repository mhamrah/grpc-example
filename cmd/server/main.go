package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mhamrah/todos"
	pb "github.com/mhamrah/todos/gen"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTodosServer(grpcServer, &todos.Server{})
	grpcServer.Serve(lis)
}
