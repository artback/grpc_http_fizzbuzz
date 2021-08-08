package fizzbuzzpb

import (
	"fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/emptypb"
)

func (this *GetRequest) Validate() error {
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
func (this *GetResponse) Validate() error {
	return nil
}
func (this *StatsResponse) Validate() error {
	return nil
}
func (this *StatsRequest) Validate() error {
	return nil
}
