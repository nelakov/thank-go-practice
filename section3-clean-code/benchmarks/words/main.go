package main

// do not remove imports, they are used by the checker
import (
	"strings"
)

type container map[string]int

// Words works with the words in a string.
type Words struct {
	words container
}

// MakeWords creates a new Words instance.
func MakeWords(s string) Words {
	wordMap := createWordsContainer(s)
	return Words{wordMap}
}

func createWordsContainer(s string) container {
	words := make(container, len(s))
	for idx, word := range strings.Fields(s) {
		if _, ok := words[word]; ok {
			continue // write only first index
		}
		words[word] = idx
	}
	return words
}

// Index returns the index of the first occurrence of the word in the string,
// or -1 if the word is not found.
func (w Words) Index(word string) int {
	if v, ok := w.words[word]; ok {
		return v
	}
	return -1
}
