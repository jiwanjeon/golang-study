/*
In this program, portal 1 sleep for 3 seconds and portal 2 sleep for 9 seconds after their sleep time over
they will ready to proceed. Now, select statement waits till their sleep time, when the portal 2 wakes up,
it selects case 2 and prints “Welcome to channel 1”. If the portal 1 wakes up before portal 2
then the output is “welcome to channel 2”.
*/

/*
In select statement, if multiple cases are ready to proceed, then one of them can be selected randomly
*/

package main

import (
	"fmt"
	"time"
)

// function 1
func portal1(channel1 chan string) {

	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}

// function 2
func portal2(channel2 chan string) {

	time.Sleep(5 * time.Second)
	channel2 <- "Welcome to channel 2"
}

// main function
func main() {

	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)

	select {

	// case 1 for portal 1
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for portal 2
	case op2 := <-R2:
		fmt.Println(op2)
	}

}
