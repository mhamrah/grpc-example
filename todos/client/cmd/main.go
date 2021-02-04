package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"strings"
	"time"

	pb "github.com/mhamrah/grpc-example/gen"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tjarratt/babble"
	"google.golang.org/grpc"
)

func setupViper() *viper.Viper {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	flag.String("backend", "localhost:50052", "The gRPC Todo endpoint to call")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return viper.GetViper()
}

func main() {

	cfg := setupViper()
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 5

	log.Printf("contacting backend: %v", cfg.GetString("backend"))
	conn, err := grpc.Dial(cfg.GetString("backend"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err: %v, could not connect to backend at %v.", err, cfg.GetString("backend"))
	}
	defer conn.Close()

	client := pb.NewTodosClient(conn)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		millis := 500 + r.Intn(500)
		<-time.After(time.Duration(millis) * time.Millisecond)

		resp, err := client.CreateTodo(context.Background(), &pb.CreateTodoRequest{Todo: &pb.Todo{Title: babbler.Babble()}})
		if err != nil {
			log.Printf("error creating todo: %v\n", err)
			continue
		}

		result, err := client.GetTodo(context.Background(), &pb.GetTodoRequest{Id: resp.Id})
		if err != nil {
			log.Printf("error creating todo: %v\n", err)
			continue
		}
		log.Printf("result id: %v\n", result.Id)
	}
}
