package main

import (
	"Test/pkg/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := api.NewMethodClient(cc)
	request := &api.AddRequest{
		X: 5,
		Y: 2,
	}

	resp, _ := client.Add(context.Background(), request)
	fmt.Println(resp.Result)

}
