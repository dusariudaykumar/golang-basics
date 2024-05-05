package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case")
	// rand is a math function which gerenerates random number
	dices := rand.Intn(6) + 1

	fmt.Println("value ", dices)

	switch dices {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		// fallthrough allows you to execute following next case as well i.e if switch falls into case 3 it will also execute next case i.e case 4
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot")
	default:
		fmt.Println("Wrong number")
	}

}
