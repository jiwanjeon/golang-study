package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	// The control does not wait for Goroutine to complete their execution just like normal function
	// they always move forward to the next line after the Goroutine call and ignores the value returned by the Goroutine.
	// So, if I add below code, the Goroutine works what I expected
	// "fork & join model concept"
	time.Sleep(time.Second)
	fmt.Println("elapsed:", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2")
}

func task3() {
	fmt.Println("task3")
}

func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4")
}

// below code change sleep into channel(improve performance)

// func main() {
// 	done := make(chan struct{})
// 	now := time.Now()
// 	go task1(done)
// 	go task2(done)
// 	go task3(done)
// 	go task4(done)

// 	<- done
// 	<- done
// 	<- done
// 	<- done
// 	fmt.Println("elapsed:", time.Since(now))
// }

// func task1(done chan struct{}) {
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("task1")
// 	done <- struct{}{}
// }

// func task2(done chan struct{}) {
// 	time.Sleep(200 * time.Millisecond)
// 	fmt.Println("task2")
// 	done <- struct{}{}
// }

// func task3(done chan struct{}) {
// 	fmt.Println("task3")
// 	done <- struct{}{}
// }

// func task4(done chan struct{}) {
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("task4")
// 	done <- struct{}{}
// }