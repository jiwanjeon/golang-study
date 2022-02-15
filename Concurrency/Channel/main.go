// Go program to illustrate
// how to create a channel
package main

import "fmt"

// How to create a channel

func main() {

	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)
}

// Send and Receive Data From a Channel

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}

// Closing a Channel

/*
close() or ele, ok:= <- Mychannel
*/
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

// Main function
func main() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
