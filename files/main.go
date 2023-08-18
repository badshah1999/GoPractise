package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hi Badshah"
	file, err := os.Create("./smp.txt")
	checkNilErr(err)
	length, er := io.WriteString(file, content)
	checkNilErr(er)
	fmt.Println("The length of file is:", length)
	defer file.Close()
	readFile("./smp.txt")
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func readFile(filepath string) {
	data, err := os.ReadFile(filepath)
	checkNilErr(err)
	fmt.Println("The file is:", string(data))
}
