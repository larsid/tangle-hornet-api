package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/larsid/tangle-hornet-api/config"
	"github.com/larsid/tangle-hornet-api/router"
)

const CONFIG_FILE_NAME = "tangle-hornet.conf"

func main() {
	port := config.GetApiPort(CONFIG_FILE_NAME, true)

	fmt.Printf("Starting server on port %s.\n", port)

	router := router.Routes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
