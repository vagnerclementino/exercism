//Package hamming allow calcuate the hamming distance
package hamming

import (
	"errors"
)

const (
	errorDNASequenceNotEqual             = "The DNA sequences must have equals length"
	errorNumberChunkIsZero               = "The chunk size cannot be zero"
	errorNumberOfChunksIsGreaterThanText = "The number of chuncks cannot be greater than text do divide"
)

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {

	var numberOfRoutines = 2

	if len(a) != len(b) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}
	if len(a) > numberOfRoutines {
		return calculateDistinctCharWithConcurrency(a, b, numberOfRoutines)
	}
	return calculateDistinctChar(a, b)
}

// calculateDistinctChar returns the number of distinct characters between two
// strings
func calculateDistinctChar(a, b string) (int, error) {

	var distinctChar = 0

	if len(a) != len(b) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distinctChar++
		}
	}
	return distinctChar, nil
}

// calculate the hamming in a concurrency way
func calculateDistinctCharWithConcurrency(a, b string, numberOfRoutines int) (int, error) {

	var distinctChar = 0

	distanceChan := make(chan int)

	chunksFromA, err := divideInChunks(a, numberOfRoutines)
	if err != nil {
		return 0, err
	}

	chunksFromB, err := divideInChunks(b, numberOfRoutines)
	if err != nil {
		return 0, err
	}

	for i := 0; i < numberOfRoutines; i++ {
		go calculateDistinctCharByChan(chunksFromA[i], chunksFromB[i], distanceChan)
	}

	for i := 0; i < numberOfRoutines; i++ {
		distinctChar += <-distanceChan
	}
	return distinctChar, nil
}

// Calculate number of distinct char using a channel
func calculateDistinctCharByChan(a, b string, distanceChan chan int) error {

	distance, err := calculateDistinctChar(a, b)
	if err == nil {
		distanceChan <- distance
	}
	return err

}

// divide a text in a list of string (chunk) with size equal the
// numberOfChunks. Ex divideInChunks("ab", 2) = ["a", "b"]
func divideInChunks(textToDivide string, numberOfChunks int) ([]string, error) {

	var chunks []string

	if textToDivide != "" {

		if err := validateDivideInChunks(textToDivide, numberOfChunks); err != nil {
			return nil, err
		}
		chunkSize := len(textToDivide) / numberOfChunks

		for i := 0; i < len(textToDivide); i += chunkSize {
			chunks = append(chunks, textToDivide[i:(i+chunkSize)])
			if len(chunks)+1 == numberOfChunks {
				chunks = append(chunks, textToDivide[i+chunkSize:])
				break
			}
		}
	}
	return chunks, nil
}

func validateDivideInChunks(textToDivide string, numberOfChunks int) error {

	if numberOfChunks == 0 {
		return errors.New(errorNumberChunkIsZero)
	}

	if textToDivide != "" && numberOfChunks > len(textToDivide) {
		return errors.New(errorNumberOfChunksIsGreaterThanText)
	}
	return nil
}
