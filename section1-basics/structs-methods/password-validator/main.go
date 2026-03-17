// Exercise: Password Validator
//
// Write a program that validates a password using a custom type and method.
// A password is valid if it satisfies both conditions:
//   1. Contains at least one letter AND at least one digit.
//   2. Has length >= 10.
//
// Do not change function signatures, the "password" type definition,
// or the "validator" variable in main().
//
// Input:  a single string with no spaces.
// Output: "true" if the password is valid, "false" otherwise.
//
// Sample Input:  helloworld1d
// Sample Output: true

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// validator checks a string against some condition
// and returns the result
type validator func(s string) bool

// digits returns true if s contains at least one digit
// according to unicode.IsDigit(), false otherwise
func digits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// letters returns true if s contains at least one letter
// according to unicode.IsLetter(), false otherwise
func letters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// minlen returns a validator that checks whether the string length
// according to utf8.RuneCountInString() is not less than the specified value
func minlen(length int) validator {
	return func(s string) bool {
		if utf8.RuneCountInString(s) < length {
			return false
		}
		return true
	}
}

// and returns a validator that accepts a string and checks
// that all funcs returned true for that string
func and(funcs ...validator) validator {
	return func(s string) bool {
		for _, f := range funcs {
			if !f(s) {
				return false
			}
		}
		return true
	}
}

// or returns a validator that accepts a string and checks
// that at least one of funcs returned true for that string
func or(funcs ...validator) validator {
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) {
				return true
			}
		}
		return false
	}
}

// password holds the password string value and a validator
type password struct {
	value string
	validator
}

// isValid checks whether the password is valid
// according to the password's assigned validator
func (p *password) isValid() bool {
	return p.validator(p.value)
}

// ┌──────────────────────────────────────┐
// │ do not change the code below         │
// └──────────────────────────────────────┘

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	// validator that checks the password contains letters and digits,
	// or its length is at least 10 characters
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
