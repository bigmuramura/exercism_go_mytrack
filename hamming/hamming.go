package hamming

import (
	"errors"
)

// Distance ... String Comparison
func Distance(a, b string) (int, error) {

	// Slice the string as a rune type
	ar, br := []rune(a), []rune(b)

	// Checking the number of slices
	if len(ar) != len(br) {
		return 0, errors.New("the number of characters is different")
	}

	ans := 0
	for i := range ar {
		if ar[i] != br[i] {
			ans++
		}
	}
	return ans, nil
}
