package hamming

import (
	"errors"
)

const (
	errorDNASequenceNotEqual             = "The DNA sequences must have equals length"
	errorChunkSizeZero                   = "The chunk size cannot be zero"
	errorNumberOfChunksIsGreaterThanText = "The number of chuncks cannot be greater than text do divide"
)

// Distance allows calculate the hamming distance
func Distance(a, b string) (int, error) {
	var numberOfRoutines = 3
	var distance = 0

	if len(a) != len(b) {
		return 0, errors.New(errorDNASequenceNotEqual)
	}

	if len(a) == 0 {
		return 0, nil
	}

	if len(a) > numberOfRoutines {

		distanceChan := make(chan int)

		chunksA, err := divideInChunks(a, numberOfRoutines)
		if err != nil {
			return 0, err
		}

		chunksB, err := divideInChunks(b, numberOfRoutines)
		if err != nil {
			return 0, err
		}

		for i := 0; i < numberOfRoutines; i++ {
			go hammingDistanceWithConcurrency(chunksA[i], chunksB[i], distanceChan)
		}

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

func divideInChunks(textToDivide string, numberOfChunks int) ([]string, error) {

	var chunks []string

	if textToDivide == "" {
		return []string{}, nil
	}

	if numberOfChunks == 0 {
		return nil, errors.New(errorChunkSizeZero)
	}

	if textToDivide != "" && numberOfChunks > len(textToDivide) {
		return nil, errors.New(errorNumberOfChunksIsGreaterThanText)
	}

	chunkSize := len(textToDivide) / numberOfChunks

	for i := 0; i < len(textToDivide); i += chunkSize {
		chunks = append(chunks, textToDivide[i:(i+chunkSize)])
		if len(chunks)+1 == numberOfChunks {
			chunks = append(chunks, textToDivide[i+chunkSize:])
			break
		}
	}
	return chunks, nil
}
