package main

import (
	"fmt"
	"net/url"
)

const URL = "https://optical.toys/privacy?hello=value"

func main() {
	fmt.Println("Handling urls")
	resp, _ := url.Parse(URL)
	fmt.Println(resp.Host)
	fmt.Println(resp.Path)
	fmt.Println(resp.RawFragment)
	fmt.Println(resp.RawQuery)
	fmt.Println(resp.Scheme)
	fmt.Printf("URL Values :\n %+v", *resp)
	fmt.Println("Query params: ", resp.Query())
}
