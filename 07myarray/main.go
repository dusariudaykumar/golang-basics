package main

import "fmt"

func main()  {
	fmt.Println("Array")

	// Array declaration 

	var fruitList  [4] string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruits list is: ",fruitList)
	fmt.Println("Fruits lenths is: ",len(fruitList))

	// Another way 

	var vegList  = [3]string{"Potato","Beens","Mushroom"}
	fmt.Println("Vegy's list is: ",vegList)
	fmt.Println("Vegy's lenths is: ",len(vegList))
}