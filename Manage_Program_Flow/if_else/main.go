package main

import (
	"fmt"
)

/*
if  var declaration;  condition {
	// code to be executed if condition is true
}
*/

func main() {
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
}
