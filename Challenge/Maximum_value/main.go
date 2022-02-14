package main

import (
	"fmt"
)

func main() {
	numArr := []int{16, 8, 42, 4, 23, 15}
	maxNum := numArr[0]

	for i := 1; i < len(numArr); i++ {
		if numArr[i] > maxNum {
			maxNum = numArr[i]
		}
	}

	max := numArr[0]

	for _, value := range numArr[1:] {
		if value > max {
			max = value
		}
	}

	fmt.Println(maxNum)
	fmt.Println(max)
}
