package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")
	datas := []string{"Apple", "Mango", "Banana"}
	for i := range datas {
		fmt.Println("Fruits", datas[i])
		if i == 1 {
			goto add
		}
	}
	for index, value := range datas {
		fmt.Printf("Index %v of fruit %v\n", index, value)
	}

add:
	for _, value := range datas {
		fmt.Printf("Values are %v\n", value)
	}
}
