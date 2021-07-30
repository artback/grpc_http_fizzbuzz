package memory

import (
	"context"
	"fmt"
)

type statistics struct {
	values map[interface{}]uint64
}

func NewMemoryStatistics() *statistics {
	return &statistics{values: make(map[interface{}]uint64)}
}

type top struct {
	frequency uint64
	value     interface{}
}

func (s statistics) GetMostUsed(_ context.Context) (interface{}, uint64, error) {
	if !(len(s.values) > 0) {
		return nil, 0, fmt.Errorf("statistics is empty")
	}
	var highest top
	for k, v := range s.values {
		if v > highest.frequency {
			highest = top{frequency: v, value: k}
		}
	}
	return highest.value, highest.frequency, nil
}

func (s statistics) UpdateStats(_ context.Context, key interface{}) {
	s.values[key]++
}
