/*
A even-ended number is a number with the same first and last digits
(1, 11, 121, etc.)

How many even-ended numbers result from multiplying two four digits umbers?
(1001 * 1011 = 1012011 is even-ended)
*/

package main

import (
	"fmt"
)

func main() {
	count := 0

	// for every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b

			// if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
