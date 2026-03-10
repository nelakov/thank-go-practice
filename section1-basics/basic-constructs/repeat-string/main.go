/*
Task: Repeat String

The program receives a string `source` and a number `times` as input.
Concatenate `source` with itself `times` times and return the result:

  source = x, times = 3 → xxx
  source = omm, times = 2 → ommomm

Sample Input:  a 5
Sample Output: aaaaa
*/

package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// the values are guaranteed to be valid
	fmt.Scan(&source, &times)

	// take the string `source` and repeat it `times` times
	// store the result in `result`
	var result string
	for i := 0; i < times; i++ {
		result += source
	}

	fmt.Println(result)
}
