// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

const DEFAULT_NAME = "you"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	return createSharedMessage(name)
}

func createSharedMessage(name string) string {
	if name == "" {
		return formatSharedMessage(DEFAULT_NAME)
	}
	return formatSharedMessage(name)
}

func formatSharedMessage(value string) string {
	return fmt.Sprintf("One for %s, one for me.", value)
}
