package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go mod verify : verifies the modules
// go mod tidy : remove's unused packages and get the required one's
// go list -m all : lists all the  installed modules
// go mod edit : to edit mod file : -go using this we can change go version ,

func main() {
	fmt.Println("About Modules")

	r := mux.NewRouter()
	r.HandleFunc("/", greater)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}

}

func greater(r http.ResponseWriter, w *http.Request) {
	r.Write([]byte("<h1>Welcome to Golang </h1>"))
}
