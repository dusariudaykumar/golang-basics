package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Catalog struct {
	CatalogId    string    `json:"catalogid"`
	ProductTitle string    `json:"title"`
	Cost         int       `json:"cost"`
	Model        string    `json:"model"`
	Variants     *Variants `json:"variants"`
}

type Variants struct {
	Color string   `json:"color"`
	Sizes []string `json:"sizes"`
}

// helper,middleware
func (c *Catalog) IsEmpty() bool {
	// return c.CatalogId == "" && c.ProductTitle == "" && c.Model == ""
	return c.ProductTitle == ""
}

// fake DB
var CatalogData []Catalog

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	// seeding
	CatalogData = append(CatalogData, Catalog{CatalogId: "1", ProductTitle: "Bella + Canvas 3001", Cost: 4, Model: "3001", Variants: &Variants{Color: "white", Sizes: []string{"S", "M", "L", "XL"}}})

	CatalogData = append(CatalogData, Catalog{CatalogId: "2", ProductTitle: "Bella + Canvas 3501", Cost: 5, Model: "3501", Variants: &Variants{Color: "Red", Sizes: []string{"S", "L", "XL"}}})

	// routes

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/catalogs", GetAllProducts).Methods("GET")
	r.HandleFunc("/catalog/{id}", GetSingleCatalogProduct).Methods("GET")
	r.HandleFunc("/catalog/{id}", UpdateOneCatalog).Methods("PUT")
	r.HandleFunc("/catalog", CreateOneCatalogProduct).Methods("POST")
	r.HandleFunc("/catalog/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// controllers

// serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang API</h1>"))
}

// get all catalog products

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all product")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CatalogData)
}

func GetSingleCatalogProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one catalog product")
	w.Header().Set("Content-Type", "application/json")

	// garb id from params
	params := mux.Vars(r)
	fmt.Println("Params: ", params)

	// loop through catalogData, find the matching id and return the response
	for _, v := range CatalogData {
		if v.CatalogId == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("No product found")
	return
}

// create a catalog product

func CreateOneCatalogProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create product")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}

	var catalog Catalog

	_ = json.NewDecoder(r.Body).Decode(&catalog)

	if catalog.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	catalog.CatalogId = strconv.Itoa(rand.Intn(1000))
	CatalogData = append(CatalogData, catalog)
	json.NewEncoder(w).Encode(catalog)
	return

}

func UpdateOneCatalog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update product")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, v := range CatalogData {
		if v.CatalogId == params["id"] {
			CatalogData = append(CatalogData[:i], CatalogData[i+1:]...)
			var updatedProduct Catalog
			json.NewDecoder(r.Body).Decode(&updatedProduct)
			updatedProduct.CatalogId = params["id"]
			CatalogData = append(CatalogData, updatedProduct)
			_ = json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}
	json.NewEncoder(w).Encode("Product not found")
	return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete API")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// find the product and delete it
	for i, v := range CatalogData {
		if v.CatalogId == params["id"] {
			CatalogData = append(CatalogData[:i], CatalogData[i+1:]...)
			json.NewEncoder(w).Encode("Successfully deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Product not found")
	return

}
