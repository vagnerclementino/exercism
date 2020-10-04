package raindrops

import "fmt"

// Convert a number into a string that contains raindrop sounds corresponding
// to certain potential factors
func Convert(input int) string {
	var result string
	var isFactor bool

	if input%3 == 0 {
		result = "Pling"
		isFactor = true
	}

	if input%5 == 0 {
		result = result + "Plang"
		isFactor = true
	}

	if input%7 == 0 {
		result = result + "Plong"
		isFactor = true
	}

	if !isFactor {
		result = fmt.Sprintf("%d", input)
	}
	return result
}
