package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dusariudaykumar/mongoapi/routes"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started")
	r := routes.Router()

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening to port 8000")
}
