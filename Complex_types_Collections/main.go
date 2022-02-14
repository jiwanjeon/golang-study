package main

import "fmt"

func main() {
	var ary = []int16{12, 7, 4, 67, 82}
	ary = append(ary[2:len(ary)])
	fmt.Println(ary)
	fmt.Println(ary[2])
}
