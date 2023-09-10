package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Index   string `json:"index"`
	Content string `json:"content"`
}

// Get all messages using a specific index.
func GetAllMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	index := vars["index"]

	fmt.Fprintf(writer, "Retornando todas as mensagens do índice: "+index)
}

// Create and submit a new message.
func CreateNewMessage(writer http.ResponseWriter, request *http.Request) {
	var message Message
	
	requestBody, _ := io.ReadAll(request.Body)

	json.Unmarshal(requestBody, &message)
	json.NewEncoder(writer).Encode(message)
}
