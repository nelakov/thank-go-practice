// 🤔 Duplicator
//
// I wrote a Produce function that duplicates a given value:
//
//   func Produce(val int, n int) []int
//
// Unfortunately, it only works for integers:
//
//   intSlice := Produce(5, 3)   // [5 5 5] — works
//   strSlice := Produce("o", 5) // [o o o o o] — doesn't work
//
// Rewrite Produce so that it works with values of any type.
//
// Submit only the code between "solution start" and "solution end" markers.
//
// Sample Output:
//
//   PASS

package main

import "fmt"

// solution start

// Produce returns a slice of n copies of val.
func Produce[T any](val T, n int) []T {
	vals := make([]T, n)
	for i := range n {
		vals[i] = val
	}
	return vals
}

// solution end

func main() {
	// works
	intSlice := Produce(5, 3)
	fmt.Println(intSlice)
	// [5 5 5]

	// doesn't work yet
	strSlice := Produce("o", 5)
	fmt.Println(strSlice)
	// [o o o o o]
}
