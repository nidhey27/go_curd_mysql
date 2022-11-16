package main

import (
	"github.com/gorilla/mux"
	"github.com/nidhey27/go-bookstore/pkg/routes"
	"log"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Println("Server Started at PORT 9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}
