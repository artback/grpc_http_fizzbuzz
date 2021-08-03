package fizzbuzzsrv

import (
	"context"
	"fmt"
	"github.com/artback/grpc_http_fizzbuzz/pkg/fizz"
	"github.com/artback/grpc_http_fizzbuzz/pkg/stats"
	"github.com/artback/grpc_http_fizzbuzz/proto/v1/fizzbuzzpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	fizzbuzzpb.UnimplementedFizzBuzzServiceServer
	stats.Statistics
}

func (s Server) Get(ctx context.Context, req *fizzbuzzpb.GetRequest) (*fizzbuzzpb.GetResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Errorf("Controller.Get Validation %w", err).Error())
	}

	buzzValues := fizz.BuzzValues{Int1: req.Int1, Int2: req.Int2, Limit: req.Limit, Str1: req.Str1, Str2: req.Str2}
	go s.UpdateStats(ctx, buzzValues)
	return &fizzbuzzpb.GetResponse{Words: fizz.RunFizz(buzzValues)}, nil
}

func (s Server) Stats(ctx context.Context, _ *fizzbuzzpb.StatsRequest) (*fizzbuzzpb.StatsResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	most, frequency, err := s.GetMostUsed(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("service.Stats: %w", err).Error())
	}
	b, ok := most.(fizz.BuzzValues)
	if !ok {
		return nil, status.Error(codes.Internal, fmt.Errorf("service.Stats: type assertion fizz.BuzzValues not ok").Error())
	}
	return &fizzbuzzpb.StatsResponse{Int1: b.Int1, Int2: b.Int2, Str1: b.Str1, Str2: b.Str2, Limit: b.Limit, Requests: frequency}, nil
}
