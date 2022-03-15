package main

import (
	"Test/pkg/api"
	"Test/pkg/method"
	"google.golang.org/grpc"
	"net"
	//"test/pkg/api"
	//"test/pkg/method"
)

func main() {
	s := grpc.NewServer()
	srv := &method.GRPCServer{}
	api.RegisterMethodServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	s.Serve(l)
}
