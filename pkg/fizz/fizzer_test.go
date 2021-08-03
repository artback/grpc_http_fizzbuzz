package fizz

import (
	"reflect"
	"testing"
)

func TestRunFizz(t *testing.T) {
	tests := []struct {
		name string
		args BuzzValues
		want []string
	}{
		{
			name: "fizzbuzzpb 2,3,5",
			args: BuzzValues{Int1: 2, Int2: 3, Limit: 5, Str1: "fizz", Str2: "buzz"},
			want: []string{
				"1", "fizz", "buzz", "fizz", "5",
			},
		},
		{
			name: "fizzbuzzpb 2,3,10",
			args: BuzzValues{Int1: 2, Int2: 3, Limit: 10, Str1: "fizz", Str2: "buzz"},
			want: []string{
				"1", "fizz", "buzz", "fizz", "5", "fizzbuzzpb", "7", "fizz", "buzz", "fizz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RunFizz(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunFizz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRunFizz(b *testing.B) {
	tests := []struct {
		name string
		args BuzzValues
		want []string
	}{
		{
			name: "fizzbuzzpb 2,3,5",
			args: BuzzValues{Int1: 2, Int2: 3, Limit: 5, Str1: "fizz", Str2: "buzz"},
			want: []string{
				"1", "fizz", "buzz", "fizz", "5",
			},
		},
		{
			name: "fizzbuzzpb 2,3,10",
			args: BuzzValues{Int1: 2, Int2: 3, Limit: 10, Str1: "fizz", Str2: "buzz"},
			want: []string{
				"1", "fizz", "buzz", "fizz", "5", "fizzbuzzpb", "7", "fizz", "buzz", "fizz",
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RunFizz(tt.args)
			}
		})
	}
}

func Test_getFizzOrBuzz(t *testing.T) {
	type args struct {
		i uint64
		n BuzzValues
	}
	tests := []struct {
		name string
		args args
		want fizzState
	}{
		{
			name: "get fizz",
			args: args{i: 2, n: BuzzValues{
				Int1: 2, Int2: 5,
			}},
			want: fizz,
		},
		{
			name: "get buzz",
			args: args{i: 5, n: BuzzValues{
				Int1: 2, Int2: 5,
			}},
			want: buzz,
		},
		{
			name: "get buzz",
			args: args{i: 6, n: BuzzValues{
				Int1: 2, Int2: 3,
			}},
			want: fizzBuzz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFizzOrBuzz(tt.args.i, tt.args.n); got != tt.want {
				t.Errorf("getFizzOrBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getFizzOrBuzz(b *testing.B) {
	type args struct {
		i uint64
		n BuzzValues
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "get fizz",
			args: args{i: 2, n: BuzzValues{
				Int1: 2, Int2: 5,
			}},
		},
		{
			name: "get buzz",
			args: args{i: 5, n: BuzzValues{
				Int1: 2, Int2: 5,
			}},
		},
		{
			name: "get buzz",
			args: args{i: 6, n: BuzzValues{
				Int1: 2, Int2: 3,
			}},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getFizzOrBuzz(tt.args.i, tt.args.n)
			}
		})
	}
}
