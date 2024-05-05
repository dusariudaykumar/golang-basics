package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// no inheritance in golang; No super or parent
	uday := User{"Uday", "uday@gmail.com", true, 22}
	fmt.Println(uday)
	fmt.Printf("Details: %v\n", uday)
	fmt.Printf("Details: %+v\n", uday)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
