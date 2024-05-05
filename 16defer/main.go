package main

import "fmt"

func main() {

	/*
		Defer:
			it executes the defer code just before the function end
	*/
	// here below code will be executed at the end of the function and if there are multiple defers the defer's will be executed in LIFO (Last In First Out)
	defer fmt.Println("Hello")
	fmt.Println("World")
	myDefers()
}

func myDefers() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
