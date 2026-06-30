package main

// do not remove imports, they are used by the checker
import (
	"strings"
	"testing"
)

// use this variable in the benchmarks
var src = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// use this variable in the benchmarks
var pattern = "commodo"

// implement the benchmark for MatchContains
func BenchmarkMatchContains(b *testing.B) {
	for b.Loop() {
		MatchContains(pattern, src)
	}
}

// implement the benchmark for MatchContainsCustom
func BenchmarkMatchContainsCustom(b *testing.B) {
	for b.Loop() {
		MatchContainsCustom(pattern, src)
	}
}

// ┌────────────────────────────────────────┐
// │ do not change the code below this line │
// └────────────────────────────────────────┘

// Library
func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// Hand-written
func MatchContainsCustom(pattern string, src string) bool {
	if pattern == "" {
		return true
	}
	if len(pattern) > len(src) {
		return false
	}
	pat_len := len(pattern)
	for idx := 0; idx < len(src)-pat_len+1; idx++ {
		if src[idx:idx+pat_len] == pattern {
			return true
		}
	}
	return false
}
