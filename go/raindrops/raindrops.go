package raindrops

import "fmt"

var factors = [...]int{3, 5, 7}
var results = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert a number into a string that contains raindrop sounds corresponding
// to certain potential factors
func Convert(input int) string {
	var result string
	for _, factor := range factors {
		if input%factor == 0 {
			result = result + results[factor]
		}
	}
	if result == "" {
		result = fmt.Sprintf("%d", input)
	}
	return result
}
