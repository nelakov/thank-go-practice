/*
Task: Abbreviation

Write a program that takes a phrase as input and builds an abbreviation
from the first letters of each word:

  Today I learned → TIL
  Высшее учебное заведение → ВУЗ
  Кот обладает талантом → КОТ

If a word does not start with a letter, ignore it:

  Ар 2 Ди #2 → АД

Only whitespace characters are considered word separators.
Hyphens, slashes, and others can be ignored:

  Анна-Мария Волхонская → АВ

Sample Input:  Today I learned
Sample Output: TIL
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	phrase := readString()

	// 1. Split the phrase into words using `strings.Fields()`
	// 2. Take the first letter of each word and convert it
	//    to upper case via `unicode.ToUpper()`
	// 3. If a word does not start with a letter, ignore it
	//    check via `unicode.IsLetter()`
	// 4. Build a word from the resulting letters and store it
	//    in the `abbr` variable

	fields := strings.Fields(phrase)

	var abbr []rune
	for _, field := range fields {
		r, _ := utf8.DecodeRuneInString(field)
		if unicode.IsLetter(r) {
			abbr = append(abbr, unicode.ToUpper(r))
		}
	}
	fmt.Println(string(abbr))
}

// ┌──────────────────────────────────────┐
// │ do not modify the code below this line │
// └──────────────────────────────────────┘

// readString reads a string from `os.Stdin` and returns it
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
