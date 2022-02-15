package main

import (
	"fmt"
)

type (
	myMap1 map[int]string
	myMap2 map[string]string
)

func (m myMap1) Contains(k int) (ok bool) {
	_, ok = m[k]
	return
}

func (m myMap2) Contains(k string) (ok bool) {
	_, ok = m[k]
	return
}

func main() {
	m1 := make(myMap1)
	m2 := make(myMap2)

	m1[10] = "Hello"

	m2["Hi"] = "こんにちは"

	fmt.Println(m1.Contains(9))  // false
	fmt.Println(m1.Contains(10)) // true

	fmt.Println(m2.Contains("Hello")) // false
	fmt.Println(m2.Contains("Hi"))    // true
}
