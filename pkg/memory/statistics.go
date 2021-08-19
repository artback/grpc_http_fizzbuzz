package memory

import (
	"context"
	"github.com/artback/grpc_http_fizzbuzz/pkg/stats"
	"sync"
)

// The time complexity for UpdateStats is  O(1) while for GetMostUsed it is O(log N).
// I figure the calls to UpdateStats will be plenty more so a good performance for those calls are more important
type statistics struct {
	values map[interface{}]uint64
	sync.RWMutex
}

func NewMemoryStatistics() *statistics {
	return &statistics{values: make(map[interface{}]uint64)}
}

type top struct {
	frequency uint64
	value     interface{}
	err       error
}

func (s *statistics) GetMostUsed(ctx context.Context) (interface{}, uint64, error) {
	select {
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	default:
		return s.getMostUsed()
	}
}

func (s *statistics) getMostUsed() (interface{}, uint64, error) {
	s.RLock()
	if !(len(s.values) > 0) {
		return nil, 0, stats.StatisticsEmpty
	}
	var highest top
	for k, v := range s.values {
		if v > highest.frequency {
			highest = top{frequency: v, value: k}
		}
	}
	s.RUnlock()
	return highest.value, highest.frequency, highest.err
}

func (s *statistics) UpdateStats(ctx context.Context, key interface{}) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.updateStats(key)
		return nil
	}
}

func (s *statistics) updateStats(key interface{}) {
	s.Lock()
	s.values[key]++
	s.Unlock()
}
