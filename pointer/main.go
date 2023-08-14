package main

import "fmt"

func main() {
	var num int = 25
	var ptr = &num
	*ptr = *ptr * 2
	fmt.Println(*ptr)
}
