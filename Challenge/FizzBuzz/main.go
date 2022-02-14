/*
if the number is divisible  / by 3 --> print "fizz" / by 5 --> "buzz" /
by both "fizzbuzz"
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	inputValue := getInput("Input value is: ")
	fmt.Println(inputValue)

	if math.Mod(inputValue, 3) == 0 && math.Mod(inputValue, 5) == 0 {
		fmt.Print("fizzbuzz\n")
	} else if math.Mod(inputValue, 5) == 0 {
		fmt.Print("buzz\n")
	} else if math.Mod(inputValue, 3) == 0 { // inputValue % 5 == 0 % 항상 math 랑 사용해야하는지??
		fmt.Print("fizz\n")
	} else {
		fmt.Print("You write wrong one")
	}
}

func getInput(prompt string) float64 {
	fmt.Printf("%v: ", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}

	return float64(value)
}
