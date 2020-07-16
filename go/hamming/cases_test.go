package hamming

import (
	"errors"
)

// Source: exercism/problem-specifications
// Commit: 4119671 Hamming: Add a tests to avoid wrong recursion solution (#1450)
// Problem Specifications Version: 2.3.0

var testCases = []struct {
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{ // empty strands
		"",
		"",
		0,
		false,
	},
	{ // single letter identical strands
		"A",
		"A",
		0,
		false,
	},
	{ // single letter different strands
		"G",
		"T",
		1,
		false,
	},
	{ // long identical strands
		"GGACTGAAATCTG",
		"GGACTGAAATCTG",
		0,
		false,
	},
	{ // long different strands
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
		false,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		0,
		true,
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		0,
		true,
	},
	{ // disallow left empty strand
		"",
		"G",
		0,
		true,
	},
	{ // disallow right empty strand
		"G",
		"",
		0,
		true,
	},
}

var divideInChunksTestCases = []struct {
	title          string
	textToDivide   string
	numberOfChunks int
	expectedChunks []string
	expectedError  error
}{
	{
		title:          "should returns an empty list if text is empty",
		textToDivide:   "",
		numberOfChunks: 1,
		expectedChunks: []string{},
		expectedError:  nil,
	},

	{
		title:          "should returns an error if numberOfChunks is zero",
		textToDivide:   "abc",
		numberOfChunks: 0,
		expectedChunks: nil,
		expectedError:  errors.New("The number of chuncks cannot be zero"),
	},

	{
		title:          "should returns an error if numberOfChunks is greater than text to divide",
		textToDivide:   "a",
		numberOfChunks: 2,
		expectedChunks: nil,
		expectedError:  errors.New("The number of chuncks cannot be greater than text do divide"),
	},

	{
		title:          "should returns a list with one item when numberOfChunks is one",
		textToDivide:   "abc",
		numberOfChunks: 1,
		expectedChunks: []string{"abc"},
		expectedError:  nil,
	},

	{
		title:          "should returns a list with size equals text to divide when numberOfChunks is equals of length of textToDivide",
		textToDivide:   "abc",
		numberOfChunks: 3,
		expectedChunks: []string{"a", "b", "c"},
		expectedError:  nil,
	},

	{
		title:          "should returns list with chunks with same length when numberOfChunks and textToDivide are multiples",
		textToDivide:   "abcdef",
		numberOfChunks: 2,
		expectedChunks: []string{"abc", "def"},
		expectedError:  nil,
	},

	{
		title:          "should returns list jwith on chunks with equals one when size of text is not multiple",
		textToDivide:   "abcdefg",
		numberOfChunks: 2,
		expectedChunks: []string{"abc", "defg"},
		expectedError:  nil,
	},
}
