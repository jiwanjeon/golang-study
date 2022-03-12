package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
  // fork point
	go func() {
		// basically tell our routine that we are done executing done and exit the function.
		defer wg.Done()
		work()
	}()
	// join point
	wg.Wait()
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func work() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("printing some stuff")
}