// Average Value
//
// I implemented a Sum type that calculates a running total of added numbers.
// But turns out the requirements were different:
//
// 1. Calculate a running average, not a running sum.
// 2. Support int and float64 types.
// 3. Add must support method chaining (x.Add(1).Add(2).Add(3)).
//
// Implement an Avg type with Add and Val methods.
//
// Example:
//
//   intAvg := Avg[int]{}
//   intAvg.Add(4).Add(3).Add(2)
//   fmt.Println(intAvg.Val()) // 3
//
//   floatAvg := Avg[float64]{}
//   floatAvg.Add(4.0).Add(3.0)
//   floatAvg.Add(2.0).Add(1.0)
//   fmt.Println(floatAvg.Val()) // 2.5
//
// Submit only the code between "solution start" and "solution end" markers.
//
// Sample Output:
//
//   PASS

package main

import "fmt"

// solution start

// Avg calculates a running average.
type Avg[T int | float64] struct {
	sum   T
	count int
}

// Add recalculates the average with the new val.
func (a *Avg[T]) Add(val T) *Avg[T] {
	a.count++
	a.sum += val
	return a
}

// Val returns the current average value.
func (a *Avg[T]) Val() T {
	if a.count == 0 {
		return T(0)
	}
	return a.sum / T(a.count)
}

// solution end

func main() {
	intAvg := Avg[int]{}
	intAvg.Add(4).Add(3).Add(2)
	fmt.Println(intAvg.Val()) // 3

	floatAvg := Avg[float64]{}
	floatAvg.Add(4.0).Add(3.0)
	floatAvg.Add(2.0).Add(1.0)
	fmt.Println(floatAvg.Val()) // 2.5
}
