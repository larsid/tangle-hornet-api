package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allancapistrano/tangle-hornet-api/endpoints"
)

const PORT = 3000

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", endpoints.Root)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}

func main() {
	fmt.Printf("Starting server on port %d", PORT)

	handleRequests()
}