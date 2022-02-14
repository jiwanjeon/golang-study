// Go program to illustrate how to pass an
// array as an argument in the function
package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10, "woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

// Dog is struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	//Go language formats according to a format specifier and returns the resulting string
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
