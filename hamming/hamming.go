package hamming

import (
	"errors"
)

// Distance ... String Comparison
func Distance(a, b string) (int, error) {

	// Slice the string as a rune type
	ar, br := []rune(a), []rune(b)

	if a == "" || b == "" {
		if a != b {
			return 0, errors.New("one is blank")
		}
	} else if len(ar) != len(br) {
		return 0, errors.New("the number of characters is different")
	}

	ans := 0
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			ans++
		}
	}
	return ans, nil
}
