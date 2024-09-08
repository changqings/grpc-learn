package main

import (
	"context"
	"fmt"
	"grpc-learn/hello"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type HelloServer struct {
	hello.UnimplementedHelloServiceServer
}

func (h *HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	slog.Info("Received request", "name", req.Name, "msg", req.Message)
	return &hello.HelloReply{
		Message: req.Message + " from server",
		Name:    req.GetName(),
	}, nil
}

var (
	server_port = ":8080"
)

func main() {

	l, err := net.Listen("tcp", server_port)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	g := grpc.NewServer()
	h := HelloServer{}

	hello.RegisterHelloServiceServer(g, &h)
	fmt.Println("grpc server running on ", server_port)

	if err := g.Serve(l); err != nil {
		grpclog.Fatalf("Failed on server: %v", err)
	}

}
