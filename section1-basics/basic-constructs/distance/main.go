/*
Task: Distance Between Points

Write a program that calculates the Euclidean distance between two points
on a plane (x1, y1) and (x2, y2):

  d = sqrt((x1-x2)^2 + (y1-y2)^2)

How the exercises work

This and all following exercises in the module are checked as follows:

- a string is passed to the program via standard input (os.Stdin);
- the program processes it and prints the result to standard output (os.Stdout);
- the printed result is compared with the expected one.

You can run the program locally by passing data via echo and a pipe.
For example, if the source file is called distance.go:

  $ echo "1 1 4 5" | go run distance.go
  5

Sample Input:  1 1 4 5
Sample Output: 5
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// declare variables x1, y1, x2, y2 of type float64
	var x1, y1, x2, y2 float64

	// read numbers from os.Stdin
	// the values are guaranteed to be valid
	// do not modify this block
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		return
	}

	// calculate d using the Euclidean distance formula
	// use math.Pow(x, 2) to square a value
	// use math.Sqrt(x) to extract the square root
	x := x1 - x2
	x = math.Pow(x, 2)
	y := y1 - y2
	y = math.Pow(y, 2)
	sum := x + y
	d := math.Sqrt(sum)

	// print the result to os.Stdout
	fmt.Println(d)
}
