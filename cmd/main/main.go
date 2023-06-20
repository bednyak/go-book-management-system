package main

import (
	"fmt"
	"github.com/bednyak/go-book-management-system/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
