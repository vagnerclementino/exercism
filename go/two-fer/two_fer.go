// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer
import "fmt"
// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var shared string = ""
	const DEFAULT_NAME = "you"
	const TEMPLATE_MESSAGE = "One for %s, one for me."
	if name == "" {

		shared = fmt.Sprintf(TEMPLATE_MESSAGE, DEFAULT_NAME)
		
	}else {		
		shared = fmt.Sprintf(TEMPLATE_MESSAGE, name)
	}
	return shared
}
