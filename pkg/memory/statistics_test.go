package memory

import (
	"context"
	"fmt"
	"github.com/artback/grpc_http_fizzbuzz/pkg/fizz"
	"reflect"
	"testing"
)

func Test_statistics_GetMostUsed(t *testing.T) {
	type fields struct {
		values map[interface{}]uint64
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		want1   uint64
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{map[interface{}]uint64{
				fizz.BuzzValues{Int1: 1}: 5,
				fizz.BuzzValues{Int1: 2}: 3,
			}},
			want:  fizz.BuzzValues{Int1: 1},
			want1: 5,
		},
		{
			name:    "empty values",
			fields:  fields{map[interface{}]uint64{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := statistics{
				values: tt.fields.values,
			}
			got, got1, err := s.GetMostUsed(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMostUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMostUsed() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetMostUsed() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Benchmark_statistics_GetMostUsed(b *testing.B) {
	type fields struct {
		values map[interface{}]uint64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "2 values to search",
			fields: fields{map[interface{}]uint64{
				fizz.BuzzValues{Int1: 1}: 5,
				fizz.BuzzValues{Int1: 2}: 3,
			}},
		},
		{
			name: "500 values to search",
			fields: fields{
				values: func() map[interface{}]uint64 {
					values := make(map[interface{}]uint64)
					for i := 0; i < 500; i++ {
						values[fizz.BuzzValues{Int1: uint64(i)}] = uint64(i + 2)
					}
					fmt.Println(len(values))
					return values
				}(),
			},
		},
		{
			name: "5000 values to search",
			fields: fields{
				values: func() map[interface{}]uint64 {
					values := make(map[interface{}]uint64)
					for i := 0; i < 5000; i++ {
						values[fizz.BuzzValues{Int1: uint64(i)}] = uint64(i + 2)
					}
					return values
				}(),
			},
		},
		{
			name:   "empty values",
			fields: fields{map[interface{}]uint64{}},
		},
	}
	for _, tt := range tests {
		s := statistics{
			values: tt.fields.values,
		}
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = s.GetMostUsed(context.Background())
			}
		})
	}
}

func Test_statistics_UpdateStats(t *testing.T) {
	type fields struct {
		values map[interface{}]uint64
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[interface{}]uint64
	}{
		{
			name:   "updateStats 1 time",
			fields: fields{map[interface{}]uint64{}},
			args: args{
				key: fizz.BuzzValues{Int1: 1},
			},
			want: map[interface{}]uint64{
				fizz.BuzzValues{Int1: 1}: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := statistics{
				values: tt.fields.values,
			}
			s.UpdateStats(context.Background(), tt.args.key)
			got := s.values
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateStats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_statistics_UpdateStats(b *testing.B) {
	type fields struct {
		values map[interface{}]uint64
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[interface{}]uint64
	}{
		{
			name:   "updateStats 1 time",
			fields: fields{map[interface{}]uint64{}},
			args: args{
				key: fizz.BuzzValues{Int1: 1},
			},
			want: map[interface{}]uint64{
				fizz.BuzzValues{Int1: 1}: 1,
			},
		},
	}
	for _, tt := range tests {
		s := statistics{
			values: tt.fields.values,
		}
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.UpdateStats(context.Background(), tt.args.key)
			}
		})
	}
}

func TestNewMemoryStatistics(t *testing.T) {
	tests := []struct {
		name string
		want *statistics
	}{
		{
			name: "create new",
			want: &statistics{values: map[interface{}]uint64{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStatistics(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryStatistics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewMemoryStatistics(b *testing.B) {
	tests := []struct {
		name string
	}{
		{
			name: "create new",
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NewMemoryStatistics()
			}
		})
	}
}
