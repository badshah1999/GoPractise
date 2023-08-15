package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.NewSource(time.Now().UnixNano())

	dice := rand.Intn(6) + 1
	//fmt.Println("Random number ", dice)

	switch dice {
	case 1:
		fmt.Println("You got 1")
	case 2:
		fmt.Println("You got 2")
	case 3:
		fmt.Println("You got 3")
		fallthrough
	case 4:
		fmt.Println("You got 4")
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	default:
		fmt.Println("Try once again")
	}
}
