package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Handling Web request")
	myurl := "https://reqres.in/api/users?page=2"
	result, _ := url.Parse(myurl)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Port:", result.Port())
	values := result.Query()

	fmt.Println(values["per_page"])
	for _, val := range values {
		fmt.Println(val)
	}

	partofURL := &url.URL{
		Scheme:   "https",
		Host:     "reqres.in",
		Path:     "/api/users",
		RawQuery: "page=2",
	}
	fmt.Println("Building URL:", partofURL.String())
}
