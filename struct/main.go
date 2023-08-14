package main

import "fmt"

func main() {
	fmt.Println("Struct in GoLang")
	values := User{"Badshah", "baasha303@gmail.com", 23}
	fmt.Printf("Values are %+v\n", values)
	fmt.Printf("Name is %v\n Email is %v\n Age is %v\n", values.Name, values.Email, values.Age)
}

type User struct {
	Name  string
	Email string
	Age   int
}
