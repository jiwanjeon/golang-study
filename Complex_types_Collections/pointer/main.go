package main

// Correct Memory Initialization
// m := make(map[string]int)
// m["theAnswer"] = 42
// fmt.Println(m)

// map[theAnswer:42]

// What is the need for the pointers?

// Variables are the names given to a memory location where the actual data is stored. To access the stored data we need the address
// of that particular memory location.
// A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points
// where the memory is located and provides ways to find out the vale stored at that memory location.
// * Operator: dereferencing operator used to declare pointer variable and access the value stored in the address
// & operator: address operator used to returns the address of a variable or to access the address of a variable to a pointer

// Initialization of pointer s with
// memory address of variable a
// var s *int = &a

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value:", *pointer1)
	fmt.Println("value1 1:", value1)
}
