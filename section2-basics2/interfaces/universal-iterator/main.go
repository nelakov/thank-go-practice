// Exercise: Universal Iterator
//
// Define an `iterator` interface that can iterate over any sequence
// element by element. The interface methods should match how they
// are used in the iterate() function below.
//
// You do NOT need to implement the interface in a concrete type —
// that comes in the next exercise.
//
// Note: main() is defined behind the scenes — do not add it.
//
// Sample Input:  1 2 3 4 5
// Sample Output:
//   1
//   2
//   3
//   4
//   5

package main

import (
	"fmt"
)

// element is the interface for a sequence element
// (empty because an element can be anything).
type element interface{}

// iterator can iterate over a sequence element by element
type iterator interface {
	next() bool
	val() element
}

// iterate walks the sequence and prints each element
func iterate(it iterator) {
	for it.next() {
		curr := it.val()
		fmt.Println(curr)
	}
}

// main() is defined behind the scenes — do not add it
