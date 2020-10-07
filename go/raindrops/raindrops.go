package raindrops

import (
	"strconv"
)

// Raindrop is type calculate raindrop problem
type Raindrop struct {
	Factor  int
	Message string
}

var raindrops = [...]Raindrop{
	{
		Factor:  3,
		Message: "Pling",
	},
	{
		Factor:  5,
		Message: "Plang",
	},
	{
		Factor:  7,
		Message: "Plong",
	},
}

// Convert a number into a string that contains raindrop sounds corresponding
// to certain potential factors
func Convert(input int) string {
	var result string
	for _, r := range raindrops {
		if input%r.Factor == 0 {
			result += r.Message
		}
	}
	if result == "" {
		result = strconv.Itoa(input)
	}
	return result
}
