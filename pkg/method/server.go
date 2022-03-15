package method

import (
	"Test/pkg/api"
	"context"
)

// GRPCServer ...
type GRPCServer struct{}

func (c *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
