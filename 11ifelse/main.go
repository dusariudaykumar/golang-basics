package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	isLoggedIn := true
	if isLoggedIn {
		fmt.Println("Logged in")
	} else {
		fmt.Println("Not logged in")
	}

	// another way

	if num := 20; num > 20 {
		fmt.Printf("%v is greater than 20 \n ", num)
	} else if num == 20 {
		fmt.Printf("%v is equals to 20 \n ", num)
	} else {
		fmt.Printf("%v is less than 20 \n ", num)
	}
}
