// Go program to illustrate how to pass an
// array as an argument in the function
package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeidght: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeidght: %v\n", poodle.Breed, poodle.Weight)
}

// Dog is struct
type Dog struct {
	Breed  string
	Weight int
}
