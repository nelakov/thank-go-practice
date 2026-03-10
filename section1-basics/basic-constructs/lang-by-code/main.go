/*
Task: Language by Code

Write a program that determines the language name by its code. Rules:

  en        → English
  fr        → French
  ru or rus → Russian
  otherwise → Unknown

Sample Input:  en
Sample Output: English
*/

package main

import (
	"fmt"
)

func main() {
	var code string
	_, err := fmt.Scan(&code)
	if err != nil {
		return
	}

	// determine the full language name by its code
	// and store it in the `lang` variable
	var lang string

	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
