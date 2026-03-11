// Exercise: Filter Collection
//
// Implement a filter() function that filters a slice of integers using a predicate
// function and returns the filtered slice. The predicate is called for each element
// of the source slice. If it returns true, the element is included in the result.
// If it returns false, it is not.
//
// Read the source slice from stdin using the provided readInput() function.
// Then apply filter() with a predicate that returns true only for even numbers.
// Print the filtered slice.
//
// Input is guaranteed to contain only integers.
//
// Sample Input:  1 2 3 4 5 6
// Sample Output: [2 4 6]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	result := make([]int, 0, len(iterable))
	for _, v := range iterable {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	src := readInput()
	res := filter(func(i int) bool { return i%2 == 0 }, src)
	fmt.Println(res)
}

// ┌──────────────────────────────────────┐
// │ do not change the code below         │
// └──────────────────────────────────────┘

// readInput reads integers from os.Stdin (space-separated)
// and returns them as a slice.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
