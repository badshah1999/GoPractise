package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// without pointer
func (u User) GetUser() {
	u.Name = "Mubhi"
	fmt.Println("Changed Name:", u.Name)
}

// with pointer
func (u *User) GetUser2() {
	u.Name = "Mubhi"
	fmt.Println("Changed Name:", u.Name)
}

func main() {
	values := User{"Badshah", 23}

	//display in a struct
	fmt.Println("Values are:", values)
	fmt.Printf("The Name is %v\n and Age is %v\n ", values.Name, values.Age)

	//calling no pointer func
	values.GetUser()

	//after calling no pointer func
	fmt.Println("Values are:", values)

	//calling pointer func
	values.GetUser2()

	//after calling pointer func
	fmt.Println("Values are:", values)
}
