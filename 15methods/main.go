package main

import "fmt"

func main() {
	fmt.Println("Methods")
	uday := User{"uday", "kumar@gmail.com", false, 22}
	fmt.Printf("User info %+v\n", uday)
	uday.GetStatus()
	email := uday.GetEMail()
	fmt.Println(email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) GetEMail() string {
	return u.Email
}
