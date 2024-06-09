package router

import (
	"github.com/larsid/tangle-hornet-api/endpoints"
	"github.com/gorilla/mux"
)

// Defines and handles the API routes
func Routes() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)

	// Node Info
	router.HandleFunc("/nodeInfo", endpoints.GetNodeInfo)
	router.HandleFunc("/nodeInfo/all", endpoints.GetAllNodeInfo)

	// Messages
	router.HandleFunc("/message", endpoints.CreateNewMessage).Methods("POST")
	router.HandleFunc("/message/{index}", endpoints.GetAllMessagesByIndex)
	router.HandleFunc(
		"/message/{index}/{maxMessages:[0-9]+}",
		endpoints.GetLastHourMessagesByIndex,
	)
	router.HandleFunc("/message/messageId/{messageID}", endpoints.GetMessageByMessageId)

	return router
}
