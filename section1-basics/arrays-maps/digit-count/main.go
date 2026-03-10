/*
Task: Digit Counter

Write a program that counts how many times each digit appears in a number.
It is guaranteed that the input contains only positive integers within the
int range.

Sample Input:  12823
Sample Output: 1:1 2:2 3:1 8:1
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return
	}

	// count how many times each digit appears in `number`
	// store the result in the `counter` map,
	// where the key is the digit and the value is
	// how many times it appears
	counter := make(map[int]int)
	for i := number; i > 0; i /= 10 {
		digit := i % 10
		counter[digit]++ // by default zero value in Go
	}

	printCounter(counter)
}

// ┌──────────────────────────────────────┐
// │ do not modify the code below this line │
// └──────────────────────────────────────┘

// printCounter prints the map in the format
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
