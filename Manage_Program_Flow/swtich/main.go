package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow { //same as switch dow := rand.Intn(7) + 1; dow {}
	case 1:
		result = "It's Sunday!"
		// fallthrough :  we can use to transfer the program control just after the statement is executed
		//in the switch cases even if the expression does not match
	case 2:
		result = "It's Monday!"
	default:
		result = "It's some other day!"
	}

	fmt.Println(result)
}
