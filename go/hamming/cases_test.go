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
	chunkSize      int
	expectedChunks []string
	expectedError  error
}{
	{
		title:          "should returns an empty list if text is empty",
		textToDivide:   "",
		chunkSize:      1,
		expectedChunks: []string{},
		expectedError:  nil,
	},

	{
		title:          "should returns an error if chunkSize is zero",
		textToDivide:   "abc",
		chunkSize:      0,
		expectedChunks: nil,
		expectedError:  errors.New("The number of chuncks cannot be zero"),
	},

	{
		title:          "should returns an error if chunkSize is greater than text to divide",
		textToDivide:   "a",
		chunkSize:      2,
		expectedChunks: nil,
		expectedError:  errors.New("The number of chuncks cannot be greater than text do divide"),
	},

	{
		title:          "should returns a list with one item when chunkSize has the same length that textToDivide",
		textToDivide:   "abc",
		chunkSize:      3,
		expectedChunks: []string{"abc"},
		expectedError:  nil,
	},

	{
		title:          "should returns list with size equals text to divide when chunkSize equals 1",
		textToDivide:   "abc",
		chunkSize:      1,
		expectedChunks: []string{"a", "b", "c"},
		expectedError:  nil,
	},
}
