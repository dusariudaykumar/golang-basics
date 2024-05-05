package main

import "fmt"

/* if the first letter is in capital letter then it is public varaible i.e it is equalient to Public const Token string */
const Token string = "hello@token"


func main()  {

	var user_name string = "uday"
	fmt.Println(user_name)

	// To get the type of the varaible
    fmt.Printf("Variable type: %T \n",user_name)

	var is_logged_in bool = false
	fmt.Println(is_logged_in)
	fmt.Printf("Variable type: %T \n",is_logged_in)

	/* @uint8 0 - 255 */
	var small_value uint8 = 255
	fmt.Println(small_value)
	fmt.Printf("Variable type: %T \n",small_value)

	var small_float float32 = 255.455455454
	fmt.Println(small_float)
	fmt.Printf("Variable type: %T \n",small_float)


	// default values and some aliases
	
	/* default value for int is 0 , string is "" */
	var another_variable int
	fmt.Println(another_variable)
	fmt.Printf("Variable type: %T \n",another_variable)

	/* Implicit type */
	var message = "hello"
	fmt.Println(message)

	/* Another way of decalring the varaible NOTE - it is only allowed inside block */
	no_of_users := 200
	fmt.Println(no_of_users)

	fmt.Println(Token)
	fmt.Printf("Variable type: %T \n",Token)
}