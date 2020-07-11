// Package twofer allows to create a 2-fer message: One for you and one for me.
package twofer

import "fmt"

const defaultName = "you"

// ShareWith receives a name and returns the message "One for X, one for me".
// The X is "you" when name is missing.
func ShareWith(name string) string {
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
