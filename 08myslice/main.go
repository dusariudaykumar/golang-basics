package main

import (
	"fmt"
	"sort"
)


func main()  {
	fmt.Println("Slices ")

	// Syntax 
	var fruitList = []string {"Apple","Mango","Peach"}

	fmt.Printf("Type of the fruitsList %T\n",fruitList) 
	fmt.Println(fruitList)

	// append 
	fruitList = append(fruitList,"Banana")
	fmt.Println(fruitList)

	// slice 
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

    fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// other ways of creating slice

	highScores := make([]int,3)

	highScores[0] = 222
	highScores[1] = 332
	highScores[2] = 899

	fmt.Println(highScores)

   // below line will throw err
	// highScores[3] = 232

	highScores = append(highScores,677,990)
	fmt.Println(highScores)

	// sorting

	sort.Ints(highScores)
	fmt.Println(highScores)
	
	fmt.Println("Is sorted: ", sort.IntsAreSorted(highScores))


	// removing items from slices based on index

	 courses := []string {"React","Ruby","TS","JS"}
	fmt.Println(courses)

    // remove ruby from courses slice
	index := 1

	// append is also used to remove items from slices
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}