package memory

import (
	"context"
	"fmt"
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
}

func (s *statistics) GetMostUsed(_ context.Context) (interface{}, uint64, error) {
	s.RLock()
	if !(len(s.values) > 0) {
		return nil, 0, fmt.Errorf("statistics is empty")
	}
	var highest top
	for k, v := range s.values {
		if v > highest.frequency {
			highest = top{frequency: v, value: k}
		}
	}
	s.RUnlock()
	return highest.value, highest.frequency, nil
}

func (s *statistics) UpdateStats(_ context.Context, key interface{}) {
	s.Lock()
	s.values[key]++
	s.Unlock()
}
