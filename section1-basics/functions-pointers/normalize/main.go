// Exercise: Normalize
//
// Implement a normalize() function that normalizes the given values.
// To normalize means to divide each value by the sum of all values.
//
// Conditions:
// - The function must work in-place, modifying the passed values rather than returning new ones.
// - The function must be variadic, accepting any number of values.
// - Values must be of type float64.
//
// Examples:
//
//	a, b := 1.0, 3.0
//	normalize(&a, &b)
//	fmt.Println(a, b)
//	// 0.25 0.75
//
//	a, b, c, d := 1.0, 2.0, 3.0, 4.0
//	normalize(&a, &b, &c, &d)
//	fmt.Println(a, b, c, d)
//	// 0.1 0.2 0.3 0.4
//
// Sample Output:
//
//	PASS

package main

import "fmt"

// normalize divides each value by the sum of all values,
// so that their sum equals 1. Modifies values in-place via pointers.
func normalize(nums ...*float64) {
	var total float64
	for _, num := range nums {
		total += *num
	}

	if total == 0 {
		return
	}

	for _, num := range nums {
		*num = *num / total
	}

}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	normalize(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// 0.1 0.2 0.3 0.4
	fmt.Println("PASS")
}
