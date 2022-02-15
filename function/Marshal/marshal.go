package main

import (
	"encoding/json"
	"fmt"
)

// declaring a struct
type Human struct {

	// defining struct variables
	Name    string
	Age     int
	Address string
}

// main function
func main() {

	// defining a struct instance
	human1 := Human{"Ankit", 23, "New Delhi"}

	// encoding human1 struct
	// into json format
	human_enc, err := json.Marshal(human1)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// as human_enc is in a byte array
	// format, it needs to be
	// converted into a string
	fmt.Println(string(human_enc))

	// converting slices from
	// golang to JSON format

	// defining an array
	// of struct instance
	human2 := []Human{
		{Name: "Rahul", Age: 23, Address: "New Delhi"},
		{Name: "Priyanshi", Age: 20, Address: "Pune"},
		{Name: "Shivam", Age: 24, Address: "Bangalore"},
	}

	// encoding into JSON format
	human2_enc, err := json.Marshal(human2)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// printing encoded array
	fmt.Println()
	fmt.Println(string(human2_enc))
}
