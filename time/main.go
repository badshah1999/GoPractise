package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")
	resp := time.Now()
	fmt.Println(resp.Format("02-01-2006"))
	fmt.Println(resp.Month())
}
