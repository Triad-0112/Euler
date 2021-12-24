package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i, j := range word {
		if unicode.IsLetter(j) && strings.ContainsRune(word[i+1:], j) {
			return false
		}
	}
	return true
}
