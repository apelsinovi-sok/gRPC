package main

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	"net"
	"test/pkg/api"
	"test/pkg/method"
)

func main() {
	s := grpc.NewServer()
	srv := &method.GRPCServer{}
	api.RegisterMethodServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	s.Serve(l)
}
