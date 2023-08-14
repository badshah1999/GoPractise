package main

import "fmt"

func main() {
	lang := make(map[string]string)
	lang["JS"] = "Javascript"
	lang["JV"] = "Java"
	lang["RB"] = "Ruby"
	fmt.Println(lang)
	//deleting a value

	delete(lang, "RB")
	fmt.Println(lang)

	for key, value := range lang {
		fmt.Printf("key %v value is %v\n", key, value)
	}
}
