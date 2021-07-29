package fizzbuzz

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/emptypb"
	"math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FizzBuzzServiceGetRequest) Validate() error {
	if this.Int1 <= 0 {
		return fmt.Errorf("validation error: FizzBuzzServiceGetRequest.Int1 must be greater than 0")
	}
	if this.Int2 <= 0 {
		return fmt.Errorf("validation error: FizzBuzzServiceGetRequest.Int2 must be greater than 0")
	}
	if this.Int1 == this.Int2 {
		return fmt.Errorf("validation error: FizzBuzzServiceGetRequest.Int1 and FizzBuzzServiceGetRequest.Int2 has the same value")
	}
	if this.Str1 == "" {
		return fmt.Errorf("validation error: FizzBuzzServiceGetRequest.str1 cannot be empty")
	}
	if this.Str2 == "" {
		return fmt.Errorf("validation error: FizzBuzzServiceGetRequest.str2 cannot be empty")
	}
	return nil
}
func (this *FizzBuzzServiceGetResponse) Validate() error {
	return nil
}
func (this *FizzBuzzServiceStatsResponse) Validate() error {
	return nil
}
func (this *FizzBuzzServiceStatsRequest) Validate() error {
	return nil
}
