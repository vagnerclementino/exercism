//Package hamming allow calcuate the hamming distance
package hamming

import (
	"errors"
)

const errorDNASequenceNotEqual = "The DNA sequences must have equals length"

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {

	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}

	var distance = 0
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}

	return distance, nil

}
