package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allancapistrano/tangle-hornet-api/router"
)

const PORT = 3000

func main() {
	fmt.Printf("Starting server on port %d", PORT)

	router := router.HandleRequests()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
