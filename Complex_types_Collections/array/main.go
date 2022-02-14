package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("Arrays: ", colors)

	var number = [5]int{5, 3, 1, 2, 4}
	fmt.Println(number)

	fmt.Println("NUmber of colors: ", len(colors))
	fmt.Println("NUmber of number: ", len(number))
}
