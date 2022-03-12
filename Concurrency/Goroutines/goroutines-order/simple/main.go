package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("go routine", i+1)
		}(i)
	}
	wg.Wait()
}
