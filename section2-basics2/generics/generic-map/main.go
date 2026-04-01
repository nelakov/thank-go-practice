// Generic Map
//
// Implement a generic Map type with the following methods:
//
// - Set: sets a value for a key.
// - Get: returns a value by key.
// - Keys: returns a slice of map keys.
// - Values: returns a slice of map values.
//
// Example:
//
//   m := Map[string, int]{}
//   m.Set("one", 1)
//   m.Set("two", 2)
//
//   fmt.Println(m.Get("one")) // 1
//   fmt.Println(m.Get("two")) // 2
//
//   fmt.Println(m.Keys())   // [one two]
//   fmt.Println(m.Values()) // [1 2]
//
// Submit only the code between "solution start" and "solution end" markers.
//
// Sample Output:
//
//   PASS

package main

import (
	"fmt"
)

// solution start

// Map is a generic key-value map.
type Map[K comparable, V any] struct {
	keys []K
	vals []V
}

// Set sets a value for a key.
func (m *Map[K, V]) Set(key K, val V) {
	for i, v := range m.keys {
		if v == key {
			m.vals[i] = val
			return
		}
	}

	m.keys = append(m.keys, key)
	m.vals = append(m.vals, val)
}

// Get returns a value by key.
func (m *Map[K, V]) Get(key K) V {
	for i, v := range m.keys {
		if v == key {
			return m.vals[i]
		}
	}

	var x V
	return x
}

// Keys returns a slice of map keys.
// Key order is unspecified and does not have to match
// the order of values from the Values method.
func (m *Map[K, V]) Keys() []K {
	return m.keys
}

// Values returns a slice of map values.
// Value order is unspecified and does not have to match
// the order of keys from the Keys method.
func (m *Map[K, V]) Values() []V {
	return m.vals
}

// solution end

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
