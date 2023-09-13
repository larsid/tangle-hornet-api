package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/allancapistrano/tangle-client-go/messages"
	"github.com/gorilla/mux"
)

type Message struct {
	Index   string `json:"index"`
	Content string `json:"content"`
}

// Get all messages using a specific index.
func GetAllMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	nodeURL := "http://127.0.0.1:14265"
	
	vars := mux.Vars(request)
	index := vars["index"]

	messagesByIndex := messages.GetAllMessagesByIndex(nodeURL, index)

	var jsonInString string
	jsonInBytes, err := json.Marshal(messagesByIndex)
	if err != nil {
		jsonInString = "{\"error\": \"Unable to convert the messages struct into JSON format.\"}"
	}

	jsonInString = string(jsonInBytes) // TODO: Corrigir espaços em branco

	fmt.Fprint(writer, jsonInString)
}

// Create and submit a new message.
func CreateNewMessage(writer http.ResponseWriter, request *http.Request) {
	var message Message
	
	requestBody, _ := io.ReadAll(request.Body)

	json.Unmarshal(requestBody, &message)
	json.NewEncoder(writer).Encode(message)
}
