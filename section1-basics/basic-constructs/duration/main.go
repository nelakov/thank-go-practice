/*
Task: Duration in Minutes

Write a program that calculates the number of minutes in a time duration.

1h30m = 90 min
300s = 5 min
10m = 10 min

Use time.Duration.Minutes() from the standard library. Throughout this course,
we will frequently use Go's standard library. It has clear documentation with
examples (https://pkg.go.dev/std) — check it out if something is unclear.

Use this and other exercises as a marker — whether the course is right for you.
The first lesson's exercises are the simplest; they will get harder. If you cannot
solve them on your own without hints, the course is probably not for you.

Solutions

After you successfully solve the exercise, a "solutions" section will open
(next to the comments) — check it to compare your solution with the author's
and other participants' solutions.

Please only publish your solution if it differs significantly from those already
published. And specify what the difference is. If the solution is weak — do not
publish it.

Sample Input:  1h30m
Sample Output: 1h30m = 90 min
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// read a time duration from os.Stdin
	// the value is guaranteed to be valid
	// do not modify this block
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	d, _ := time.ParseDuration(s)

	// print the original value
	// and the number of minutes in it
	// in the format "original = X min"
	// use the .Minutes() method of d
	fmt.Printf("%s = %.0f min", s, d.Minutes())
}
