package hamming

import (
	"errors"
)

const (
	errorDNASequenceNotEqual             = "The DNA sequences must have equals length"
	errorChunkSizeZero                   = "The chunk size cannot be zero"
	errorNumberOfChunksIsGreaterThanText = "The number of chuncks cannot be greater than text do divide"
)

// Distance ...
func Distance(a, b string) (int, error) {
	var numberOfRoutines = 2
	var distance = 0
	distanceChan := make(chan int)

	if len(a) != len(b) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}
	if len(a) == 0 {
		return 0, nil
	}

	if len(a) > 2 {
		go hammingDistanceWithConcurrency(a[:len(a)/2], b[:len(b)/2], distanceChan)
		go hammingDistanceWithConcurrency(a[len(a)/2:], b[len(b)/2:], distanceChan)
		for i := 0; i < numberOfRoutines; i++ {
			distance += <-distanceChan

		}
		return distance, nil
	}
	return hammingDistance(a, b)
}

func hammingDistanceWithConcurrency(a, b string, distanceChan chan int) error {
	distance, err := hammingDistance(a, b)
	if err != nil {
		return err
	}
	distanceChan <- distance
	return nil
}

func hammingDistance(a, b string) (int, error) {
	var distance = 0
	if len(a) != len(b) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

func divideInChunks(textToDivide string, chunkSize int) ([]string, error) {

	if textToDivide == "" {
		return []string{}, nil
	}

	if chunkSize == 0 {
		return nil, errors.New(errorChunkSizeZero)
	}

	if textToDivide != "" && chunkSize > len(textToDivide) {
		return nil, errors.New(errorNumberOfChunksIsGreaterThanText)
	}

	if len(textToDivide) == chunkSize {
		return []string{textToDivide}, nil
	}

	var chunks []string
	var res string
	for i, r := range textToDivide {
		res = res + string(r)
		if (i+1)%chunkSize == 0 {
			chunks = append(chunks, res)
			res = ""
		}
	}
	return chunks, nil
}
