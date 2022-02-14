package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}

	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
		if value > 8 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of program")
}
