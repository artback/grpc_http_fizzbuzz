package fizzbuzzsrv

import (
	"context"
	"errors"
	"github.com/artback/grpc_http_fizzbuzz/mocks"
	"github.com/artback/grpc_http_fizzbuzz/pkg/fizz"
	"github.com/artback/grpc_http_fizzbuzz/proto/v1/fizzbuzz"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestService_Get(t *testing.T) {
	type args struct {
		request   *fizzbuzz.FizzBuzzServiceGetRequest
		sendError error
	}
	type expect struct {
		arg1  interface{}
		times int
	}
	tests := []struct {
		name string
		args
		expect  expect
		want    *[]string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				request: &fizzbuzz.FizzBuzzServiceGetRequest{Int1: 2, Int2: 3, Limit: 5, Str1: "fizz", Str2: "buzz"},
			},
			expect: expect{
				arg1: fizz.BuzzValues{
					Int1: 2, Int2: 3, Limit: 5, Str1: "fizz", Str2: "buzz",
				},
				times: 1,
			},
			want: &[]string{
				"1", "fizz", "buzz", "fizz", "5",
			},
		},
		{
			name: "error validate",
			args: args{
				request: &fizzbuzz.FizzBuzzServiceGetRequest{Int1: 1, Int2: 1, Limit: 5, Str1: "fizz", Str2: "buzz"},
			},
			expect: expect{
				times: 0,
			},
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _, tt := range tests {
		stats := mocks.NewMockStatistics(ctrl)
		s := Service{Statistics: stats}
		t.Run(tt.name, func(t *testing.T) {
			stats.EXPECT().UpdateStats(gomock.Any(),
				tt.expect.arg1,
			).Times(tt.expect.times)
			got, err := s.Get(context.Background(), tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != nil {
				want := &fizzbuzz.FizzBuzzServiceGetResponse{Words: *tt.want}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("Get() got= %v, want=%v", got, want)
				}
			}
		})
	}
}

func TestService_Stats(t *testing.T) {
	type Return struct {
		v         interface{}
		frequency uint64
		err       error
	}
	tests := []struct {
		name    string
		ret     Return
		want    *fizzbuzz.FizzBuzzServiceStatsResponse
		wantErr bool
	}{
		{
			name: "success",
			want: &fizzbuzz.FizzBuzzServiceStatsResponse{
				Int1: 1, Int2: 2, Requests: 1,
			},
			ret: Return{v: fizz.BuzzValues{Int1: 1, Int2: 2}, frequency: 1, err: nil},
		},
		{
			name:    "bad type assertion",
			ret:     Return{v: Return{}, frequency: 1, err: nil},
			wantErr: true,
		},
		{
			name:    "bad type assertion",
			ret:     Return{v: fizz.BuzzValues{Int1: 1, Int2: 2}, frequency: 1, err: errors.New("something happened")},
			wantErr: true,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tt := range tests {
		stats := mocks.NewMockStatistics(mockCtrl)
		t.Run(tt.name, func(t *testing.T) {
			stats.EXPECT().GetMostUsed(gomock.Any()).Times(1).Return(tt.ret.v, tt.ret.frequency, tt.ret.err)
			s := Service{Statistics: stats}
			got, err := s.Stats(context.Background(), &fizzbuzz.FizzBuzzServiceStatsRequest{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stats() got = %v, want %v", got, tt.want)
			}
		})
	}
}
