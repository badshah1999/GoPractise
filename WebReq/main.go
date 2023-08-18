package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://reqres.in/api/users?page=2"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Response type is %T:", response)
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is:", string(data))
}
