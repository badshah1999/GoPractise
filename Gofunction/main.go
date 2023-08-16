package main

import "fmt"

func main() {
	fmt.Println("Functions in GoLang")
	greet()
	res, msg := add(5, 6)
	fmt.Println("MSG:", msg)
	fmt.Println("Answer:", res)
	addRes, msg2 := advAdd(4, 5, 6, 77, 2)
	fmt.Println("Title", msg2)
	fmt.Println("Result", addRes)
}
func greet() {
	fmt.Println("Hello Gopher")
}
func add(val1 int, val2 int) (int, string) {
	return val1 + val2, "The Addition is"
}
func advAdd(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}
	return total, "Multiple values Additon:"
}
