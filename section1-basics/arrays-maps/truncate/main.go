/*
Task: Truncate Byte String

Write a program that truncates a string to the specified length and appends
an ellipsis:

  text = Eyjafjallajokull, width = 6 → Eyjafj...

If the string does not exceed the specified length, leave it unchanged:

  text = hello, width = 6 → hello

It is guaranteed that the source string `text` contains only single-byte
characters without spaces, and `width` is strictly greater than 0.

Sample Input:  Eyjafjallajokull 6
Sample Output: Eyjafj...
*/

package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	_, err := fmt.Scanf("%s %d", &text, &width)
	if err != nil {
		return
	}

	if len(text) <= width {
		fmt.Println(text)
		return
	}

	// take the first `width` bytes of `text`,
	// append `...` and store the result in `res`
	res := text[:width] + "..."

	fmt.Println(res)
}
