package main

import (
	"fmt"
	"sort"
)

func main() {
	var array [4]int
	array[0] = 1
	array[1] = 3
	array[2] = 5
	array[3] = 9
	fmt.Println(array)
	fmt.Println(len(array))
	// array = append(array,)
	var item = []string{"Apple", "Mango", "Banana"}
	fmt.Println(item)
	item = append(item, "Pomo", "Orange", "Papaya")
	fmt.Println(item[:])

	//another method of declaring array

	scores := make([]int, 2)
	scores[0] = 1
	scores[1] = 2
	fmt.Println(scores)
	scores = append(scores, 4, 3, 590, 56)
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println(scores)
}
