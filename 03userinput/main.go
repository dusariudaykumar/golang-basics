package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)

/*  bufio is a package used to read input */
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Age: ")


	/* below syntax is called comma ok || comma err*/
	input,_ := reader.ReadString('\n')

	fmt.Println("Your age is ",input)
	fmt.Printf("Type of age is %T ",input)
}