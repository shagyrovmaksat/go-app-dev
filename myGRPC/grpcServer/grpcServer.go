package grpcServer

import (
	api "awesomeProject/proto"
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) Reverse(ctx context.Context, req *api.ReverseRequest) (*api.ReverseResponse, error) {
	return &api.ReverseResponse{Result: reverse(req.GetMessage())}, nil
}

func reverse(message string) string {
	result := ""
	for _, c := range message {
		result = string(c) + result
	}
	return result
}
