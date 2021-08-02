package fizz

import "strconv"

type BuzzValues struct {
	Str1  string
	Str2  string
	Int1  uint64
	Int2  uint64
	Limit uint64
}

func RunFizz(f BuzzValues) []string {
	words := make([]string, 0, f.Limit)
	for i := uint64(1); i <= f.Limit; i++ {
		var word string
		switch getFizzOrBuzz(i, f) {
		case fizzBuzz:
			word = f.Str1 + f.Str2
		case fizz:
			word = f.Str1
		case buzz:
			word = f.Str2
		case number:
			word = strconv.FormatUint(i, 10)
		}
		words = append(words, word)
	}
	return words
}

type fizzState uint8

const (
	fizzBuzz fizzState = iota
	fizz
	buzz
	number
)

func getFizzOrBuzz(i uint64, f BuzzValues) fizzState {
	if i%(f.Int1*f.Int2) == 0 {
		return fizzBuzz
	} else if i%f.Int1 == 0 {
		return fizz
	} else if i%f.Int2 == 0 {
		return buzz
	} else {
		return number
	}

}
