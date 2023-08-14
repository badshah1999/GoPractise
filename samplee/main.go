package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	in, _ := reader.ReadString('\n')
	fmt.Println("thanks", input)
	fmt.Println("welcome", in)
}
