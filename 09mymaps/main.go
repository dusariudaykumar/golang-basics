package main

import "fmt"

func main(){
	fmt.Println("Maps in golang: ")
	// Maps are nothing but key value pairs

	languages := make(map[string]string,3)
    languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["RB"] = "Ruby"

	fmt.Println("List of langs: ",languages)
	fmt.Println("JS is ", languages["JS"])

	// delete
	
	delete(languages,"RB")
	fmt.Println("List of langs: ",languages)

	// for loop

	for key,value := range languages{
		fmt.Printf("Key : %v, Value: %v\n",key,value)
	}
}