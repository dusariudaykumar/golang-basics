package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Time functions");

	present_time := time.Now()

	/* output :  2024-04-28 14:20:41.819473 +0530 IST m=+0.000274585 */ 
	fmt.Println("Present time: ", present_time)

	/* to formate the time */
	/* we need to use the below syntax in order to format the time, for more reference checkout the offical doc  */
	fmt.Println("Formated time: ",present_time.Format("01-02-2006 15:04:05 Monday"))

	/* to create date */
	create_date := time.Date(2024,time.May,21,21,21,21,21,time.UTC) 
	fmt.Println(create_date)
	fmt.Println("Formated time: ",create_date.Format("01-02-2006 15:04:05 Monday"))

}