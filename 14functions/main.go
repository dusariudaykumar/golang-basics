package main

import "fmt"

func main() {
	fmt.Println("Function")
	// function with no parameter
	greater()

	result := addTwoNumbers(2, 5)
	fmt.Println(result)

	proResult, proMessage := proAdder(1, 3, 2, 4, 5)
	fmt.Println(proResult)
	fmt.Println(proMessage)
}

func greater() {
	fmt.Println("Welcome ")
}

func addTwoNumbers(num1 int, num2 int) int {
	return num1 + num2

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "Hello "
}
