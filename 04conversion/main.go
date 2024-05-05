package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welocome")
	fmt.Println("Please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin);

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,  ",input)

	/* we get input as type string so, to convert we use strconv package  */
	/* to trim spaces we use strings package */
	num_rating,err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Adding 1 to you're rating: ",num_rating + 1)
	}


}