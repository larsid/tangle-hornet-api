package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/larsid/tangle-hornet-api/config"
	"github.com/larsid/tangle-hornet-api/router"
)

const CONFIG_FILE_NAME = "tangle-hornet.conf"

func main() {
	port := config.GetApiPort(CONFIG_FILE_NAME, true)

	fmt.Printf("Starting server on port %s.\n", port)

	r := router.Routes()

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      3 * time.Minute,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}
