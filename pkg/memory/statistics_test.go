package memory

import (
	"context"
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
