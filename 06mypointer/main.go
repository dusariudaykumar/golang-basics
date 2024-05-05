package main

import "fmt"

func main(){
	fmt.Println("Pointers")

	/* @pointer : sometimes when we pass a variables a copy of varaibles is passed on whenever 
	there is a case when you want to avoid such things to happen and you want 
	absolute guarantee that always actual value should be passed on then 
	we prefer that a pointer should be passed, Pointer is nothing its just direct reference to the memory address 
	 */

	// pointer creatation
    var ptr *int
	fmt.Println("Value of pointer ",ptr)

	my_num := 23

	// reference means & 
	var ptr1 = &my_num

	/* it prints memory address of the pointer */
	fmt.Println("value fo actual pointer is ",ptr1)

	/* it prints value i.e 23 */
	fmt.Println("value fo actual pointer is ",*ptr1)

	*ptr1 = *ptr1 + 2
    fmt.Println("New value is: ",my_num)
}