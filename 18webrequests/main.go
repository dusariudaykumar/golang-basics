package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://optical.toys"

func main() {
	fmt.Println("Web request")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", resp)

	defer resp.Body.Close() // connection won't be closed until we close

	dataBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body: ", string(dataBytes))

}
