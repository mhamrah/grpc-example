package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	pb "github.com/mhamrah/grpc-example/gen"
	todos "github.com/mhamrah/grpc-example/todos/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))
	if err != nil {
		sugar.With("error", err).Fatal("failed to listen on port")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTodosServer(grpcServer, todos.NewServer(todos.MemoryStorage{}, sugar.Desugar()))

	reflection.Register(grpcServer)

	signalRunner(func() {
		sugar.Info("Starting server, hit Ctrl-C to stop...")
		grpcServer.Serve(lis)
	},
		func() {
			sugar.Info("Stopping server...")
			grpcServer.GracefulStop()
			sugar.Info("Bubye Y'all!")
		})
}

func signalRunner(runner, stopper func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	go func() {
		runner()
	}()

	select {
	case <-signals:
		stopper()
	}
}
