package main

import "fmt"

func main() {
	s := []int{100, 101, 102}
	for v := range s {
		fmt.Println(v) // index
	}
	fmt.Println("----------")
	for _, v := range s {
		fmt.Println(v) // element(value)
	}
	fmt.Println("----------")
	for i, v := range s {
		fmt.Println(i, "=>", v) // index & element
	}
	fmt.Println("Hello, 世界")
}
