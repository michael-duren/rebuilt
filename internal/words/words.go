// Package words has a loader for loading 460k words
// into a []string for various use cases, auto complete,
// path testing, etc
package words

import (
	_ "embed"
	"strings"
)

//go:embed words.txt
var raw string

var words []string

func init() {
	words = strings.Split(raw, "\n")
}

// GetWords retrieves a slice of
// english words up to around 450k
// sorted lexicographically
func GetWords(length int) []string {
	if length >= len(words) {
		return words
	}

	return words[0:length]
}
