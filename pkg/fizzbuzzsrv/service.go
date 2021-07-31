package fizzbuzzsrv

import (
	"context"
	"fmt"
	"github.com/artback/grpc_http_fizzbuzz/pkg/fizz"
	"github.com/artback/grpc_http_fizzbuzz/pkg/stats"
	"github.com/artback/grpc_http_fizzbuzz/proto/v1/fizzbuzz"
)

type Service struct {
	fizzbuzz.UnimplementedFizzBuzzServiceServer
	stats.Statistics
}

func (s Service) Get(ctx context.Context, req *fizzbuzz.FizzBuzzServiceGetRequest) (*fizzbuzz.FizzBuzzServiceGetResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("Service.Get Validation %w", err)
	}

	buzzValues := fizz.BuzzValues{Int1: req.Int1, Int2: req.Int2, Limit: req.Limit, Str1: req.Str1, Str2: req.Str2}
	go s.UpdateStats(ctx, buzzValues)
	return &fizzbuzz.FizzBuzzServiceGetResponse{Words: fizz.RunFizz(buzzValues)}, nil
}

func (s Service) Stats(ctx context.Context, _ *fizzbuzz.FizzBuzzServiceStatsRequest) (*fizzbuzz.FizzBuzzServiceStatsResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	most, frequency, err := s.GetMostUsed(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.Stats: %w", err)
	}
	b, ok := most.(fizz.BuzzValues)
	if !ok {
		return nil, fmt.Errorf("service.Stats: type assertion fizz.BuzzValues not ok")
	}
	return &fizzbuzz.FizzBuzzServiceStatsResponse{Int1: b.Int1, Int2: b.Int2, Str1: b.Str1, Str2: b.Str2, Limit: b.Limit, Requests: frequency}, nil
}
