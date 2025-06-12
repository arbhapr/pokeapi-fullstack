package helper

import (
	"strings"
	"unicode"
)

// ucwords capitalizes the first letter of each word in a string
func Ucwords(s string) string {
	// Split the string into words
	words := strings.Fields(s)

	// Iterate over each word and capitalize the first letter
	for i, word := range words {
		if len(word) > 0 {
			// Capitalize the first letter and append the rest of the word
			words[i] = string(unicode.ToUpper(rune(word[0]))) + word[1:]
		}
	}

	// Join the words back into a single string
	return strings.Join(words, " ")
}
