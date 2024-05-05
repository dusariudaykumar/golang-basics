package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web requests ")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	url := "http://localhost:8080/v1/catalog/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// one way of reading respone
	content, _ := io.ReadAll(resp.Body)
	fmt.Println("Response: ", string(content))

	// other way
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println("Response: ", responseString.String())
}

func PerformPostRequest() {
	url := "http://localhost:8080/v1/catalog/post"

	//fake json paylod

	requestBody := strings.NewReader(`{
		"courseName":"React JS",
		"price":"free",
		"platform":"youtube"
	}`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response: ", string(content))
}

func PerformPostFormRequest() {
	URL := "http://localhost:8080/v1/catalog/post"

	// formdata

	data := url.Values{}
	data.Add("firstname", "Uday")
	data.Add("lastname", "kumar")

	response, err := http.PostForm(URL, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response: \n", string(content))
}
