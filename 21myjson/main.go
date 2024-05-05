package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"courseName"` // json alias is used to formate the keys
	Price    int
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"` // omitempty if tags are empty it won't to shown
	Password string   `json:"-"`              // "-" won't be shown inside json
}

func main() {
	fmt.Println("Welcome to JSON ")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []Course{
		{"React js", 29, "youtube", []string{"web dev", "react"}, "abcd"},
		{"Golang", 29.00, "youtube", []string{"web dev", "golang"}, "efgh"},
		{"ruby", 29.00, "youtube", nil, "efgh"},
	}

	formatedJson, _ := json.MarshalIndent(courses, "", "\t")
	fmt.Printf("%s\n", formatedJson)
}

func DecodeJson() {
	// any data comes from web is in byte formate
	jsonDataFromWeb := []byte(`
	 {
                "courseName": "React js",
                "Price": 29,
                "website": "youtube",
                "tags": [
                        "web dev",
                        "react"
                ]
        }
	`)
	var jsonData Course

	if json.Valid(jsonDataFromWeb) {
		fmt.Println("It is an valid json")
		json.Unmarshal(jsonDataFromWeb, &jsonData)
		fmt.Printf("%#v\n", jsonData)
	} else {
		fmt.Println("INVALID JSON")
	}

	// some cases where uou just want to add data to key value

	var myCourses map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myCourses)
	fmt.Printf("%#v\n", myCourses)

	for k, v := range myCourses {
		fmt.Printf("Key: %v, Value: %v, Type: %T\n", k, v, v)
	}
}
