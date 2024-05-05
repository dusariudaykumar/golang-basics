package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in golang")

	content := "Hello Welcom To Golang"

	file, err := os.Create("./myfile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	CheckNilErr(err)

	lenght, err := io.WriteString(file, content)

	CheckNilErr(err)

	fmt.Println(lenght)

	defer file.Close()
	readFile(file.Name())
}

func readFile(fileName string) {
	dataBytes, err := os.ReadFile(fileName)

	CheckNilErr(err)

	fmt.Println("File content: ", string(dataBytes))
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
