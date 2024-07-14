package main

import (
	"context"
	"fmt"
	"grpc-learn/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	server_port = "127.0.0.1:8080"
)

func main() {
	conn, err := grpc.NewClient(server_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc client error on:%v", err)
	}

	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)
	r, err := client.SayHello(context.Background(), &hello.HelloRequest{
		Name:    "client",
		Message: "hello",
	})
	if err != nil {
		log.Fatalf("grpc client SayHello error:%v", err)
	}
	fmt.Printf("server res: %v\n", r)

}
