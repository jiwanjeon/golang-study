package main

import (
	"fmt"
	"sort"
)

func main() {
	// with one difference you are not allowed to specify the size of the slice in the square braces[]
	var colors = []string{"Rdd", "Green", "Blue"}
	fmt.Println("Arrays: ", colors)

	colors = append(colors, "Purple")
	fmt.Println("Arrays: ", colors)

	colors = append(colors[1:len(colors)])
	fmt.Println("Arrays: ", colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println("Arrays: ", colors)

	//Using make() function for making slice: make([]T, len, cap)
	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 124
	numbers[2] = 154
	numbers[3] = 34
	numbers[4] = 13
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println("Sorting: ", numbers)

	// Creating an array
	arr := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}

	// Display array
	fmt.Println("Array:", arr)

	// Creating a slice
	myslice := arr[1:6]

	// Display slice
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice))

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))

	// How to split ==> 이름 := bytes.Split(이름, []byte("기준"))
}
