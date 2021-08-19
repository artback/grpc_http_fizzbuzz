package stats

import (
	"context"
	"errors"
)

var StatisticsEmpty = errors.New("statistics is empty")

//go:generate mockgen -destination=../../mocks/mock_statistics.go -package=mocks github.com/artback/grpc_http_fizzbuzz/pkg/stats Statistics
type Statistics interface {
	GetMostUsed(context.Context) (interface{}, uint64, error)
	UpdateStats(context.Context, interface{}) error
}
