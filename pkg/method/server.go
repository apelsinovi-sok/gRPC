package method 

import (
	"context"
	"test/pkg/api"
)


// GRPCServer ...
type GRPCServer struct{}

func (c *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &app.AddResponse{Result: req.GetX() + req.GetY()}, nil
}