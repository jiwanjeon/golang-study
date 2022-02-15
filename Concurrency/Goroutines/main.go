package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		// The control does not wait for Goroutine to complete their execution just like normal function
		// they always move forward to the next line after the Goroutine call and ignores the value returned by the Goroutine.
		// So, if I add below code, the Goroutine works what I expected
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("GeeksforGeeks")
}
