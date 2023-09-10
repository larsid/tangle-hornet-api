package router

import (
	"github.com/allancapistrano/tangle-hornet-api/endpoints"
	"github.com/gorilla/mux"
)

func Routes() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)

	// Root
	router.HandleFunc("/", endpoints.Root)

	// Messages
	router.HandleFunc("/message", endpoints.CreateNewMessage).Methods("POST")
	router.HandleFunc("/message/{index}", endpoints.GetAllMessagesByIndex)

	return router
}
