package isbnverifier

import "strings"

const ISBN_SIZE = 10

func IsValidISBN(isbn string) bool {
	result := 0
	cleanIsbn := strings.ReplaceAll(isbn, "-", "")
	if !isValidFormat(cleanIsbn) {
		return false
	}
	for idx, ch := range cleanIsbn {
		currentVal := 0
		if isCheckDigit(ch, idx) {
			currentVal = ISBN_SIZE
		} else {
			currentVal = int(ch - '0')
		}
		result += currentVal * (ISBN_SIZE - idx)
	}
	return result != 0 && result%11 == 0
}

func isValidFormat(isbn string) bool {
	return len(isbn) == 10
}

func isCheckDigit(ch int32, idx int) bool {
	return ch == 'X' && idx == ISBN_SIZE-1
}
