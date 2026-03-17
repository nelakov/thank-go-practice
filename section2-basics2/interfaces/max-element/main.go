// Exercise: Max Element of a Sequence
//
// Define the `iterator` interface that can be used in the max() function.
// Implement the interface for a slice of ints (intIterator).
//
// Details are in the code below.
//
// Sample Input:  1 4 5 2 3
// Sample Output: 5

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// element is the interface for a sequence element
type element interface{}

// weightFunc returns the weight of an element
type weightFunc func(element) int

// iterator can iterate over a sequence element by element
type iterator interface {
	next() bool
	val() element
}

// intIterator iterates over a slice of ints
// (implements the iterator interface)
type intIterator struct {
	src []int
	idx int
}

// intIterator methods that implement the iterator interface
func (i *intIterator) next() bool {
	i.idx++
	if i.idx < len(i.src) {
		return true
	}
	return false
}

func (i *intIterator) val() element {
	return i.src[i.idx]
}

// newIntIterator constructs an intIterator
func newIntIterator(src []int) *intIterator {
	return &intIterator{
		src: src,
		idx: -1,
	}
}

// ┌──────────────────────────────────────┐
// │ do not change the code below         │
// └──────────────────────────────────────┘

// main finds the maximum number from the input.
func main() {
	nums := readInput()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max returns the maximum element in the sequence.
// Elements are compared by weight returned by the weight function.
func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput reads a sequence of ints from os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
