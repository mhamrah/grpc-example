package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/mhamrah/grpc-example/gen"
	todos "github.com/mhamrah/grpc-example/todos/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTodosServer(grpcServer, todos.NewServer(todos.MemoryStorage{}))

	reflection.Register(grpcServer)

	signalRunner(func() {
		log.Println("Starting server...")
		grpcServer.Serve(lis)
	},
		func() {
			log.Println("Stopping server...")
			grpcServer.GracefulStop()
			log.Println("Bubye!")
		})
}

func signalRunner(runner, stopper func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	go func() {
		runner()
	}()

	log.Println("hit Ctrl-C to shutdown")
	select {
	case <-signals:
		stopper()
	}
}
