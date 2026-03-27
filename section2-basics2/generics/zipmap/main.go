// Docking
//
// Implement a generic function ZipMap that pairs the given keys
// and values and returns the resulting map:
//
//   keys := []string{"one", "two", "thr"}
//   vals := []int{11, 22, 33}
//
//   m := ZipMap(keys, vals)
//   fmt.Println(m)
//   // map[one:11 two:22 thr:33]
//
// If keys and vals have different lengths, only the first N
// key/value pairs should be included, where N is the minimum
// of len(keys) and len(vals):
//
//   m := ZipMap([]string{"one"}, []int{11, 22, 33})
//   // map[one:11]
//
//   m := ZipMap([]string{}, []int{11, 22, 33})
//   // map[]
//
// Submit only the code between "solution start" and "solution end" markers.
//
// Sample Output:
//
//   PASS

package main

import "fmt"

// solution start

// ZipMap returns a map where keys are elements from keys and values are from vals.
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	size := min(len(keys), len(vals))
	m := make(map[K]V, size)

	for i := 0; i < size; i++ {
		m[keys[i]] = vals[i]
	}

	return m
}

// solution end

func main() {
	m1 := ZipMap([]string{"one", "two", "thr"}, []int{11, 22, 33})
	fmt.Println(m1)
	// map[one:11 two:22 thr:33]

	m2 := ZipMap([]string{"one"}, []int{11, 22, 33})
	fmt.Println(m2)
	// map[one:11]

	m3 := ZipMap([]string{"one", "two", "thr"}, []int{11})
	fmt.Println(m3)
	// map[one:11]

	m4 := ZipMap([]string{}, []int{11, 22, 33})
	fmt.Println(m4)
	// map[]
}
