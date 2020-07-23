//Package hamming allow calcuate the hamming distance
package hamming

import (
	"errors"
)

const errorDNASequenceNotEqual = "The DNA sequences must have equals length"

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}

	var distance = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
