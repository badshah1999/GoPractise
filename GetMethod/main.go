package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Get Request")
	GetMethod("http://localhost:1111/get")
}
func GetMethod(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println("Content lenght:", res.ContentLength)
	content, _ := io.ReadAll(res.Body)
	fmt.Println("Message:", string(content))
}
