package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// type - 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// type - 2
	for i := range days {
		fmt.Println(days[i])
	}

	// type - 3

	for index, day := range days {
		fmt.Printf("Index: %v and Day: %v\n", index, day)
	}

	// type - 4 similar to while

	index := 1
	for index < 10 {
		// break
		// if index == 5 {
		// 	break
		// }

		//continue
		if index == 3 {
			index++
			continue
		}
		if index == 2 {
			goto goToLable
		}

		fmt.Println("Index: ", index)
		index++
	}

goToLable:
	fmt.Println("Hello")

}
