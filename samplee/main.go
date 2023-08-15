package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// in, _ := reader.ReadString('\n')
	// fmt.Println("thanks", input)
	// fmt.Println("welcome", in)

	var firstName string
	fmt.Println("Enter you name:")
	fmt.Scanln(&firstName)
	fmt.Println("Hi", firstName)
}
