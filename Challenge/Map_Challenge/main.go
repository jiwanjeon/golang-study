// /*
// 	Word count
// */

// package main

// import "strings"

// func main() {
// 	text := `
// 	Needles and pins
// 	Needles and pins
// 	Sew me a sail
// 	To catch me the wind
// 	`
// 	wordSplit := strings.Fields(text)

// 	wc := make(map[string]int)
// 	for _, word := range wordSplit {
// 		_, matched := wc[word]
// 		if matched {
// 			wc[word] += 1
// 		} else {
// 			wc[word] = 1
// 		}
// 	}

// 	return wc

// }

// Golang program to count the number of
// repeating words in given Golang String
package main

import (
	"fmt"
	"strings"
)

func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)

	for _, word := range input {
		fmt.Println("wc: ", wc)
		fmt.Println("word: ", word)

		// TODO: How matched print out true or false --> is it the same as if in(python) & include(javascript)
		_, matched := wc[word]
		fmt.Println("matched: ", matched)
		// The special form v, ok := someMap[key] retrieves the v value associated with key key,
		//and also reports if key is in the map.

		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

// Main function
func main() {
	input := "betty bought the butter , the butter was bitter , " +
		"betty bought more butter to make the bitter butter better "
	for index, element := range repetition(input) {
		fmt.Println(index, "=", element)
	}
}
