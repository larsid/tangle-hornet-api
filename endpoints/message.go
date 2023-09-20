package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/allancapistrano/tangle-client-go/messages"
	"github.com/allancapistrano/tangle-hornet-api/config"
	"github.com/gorilla/mux"
)

type Message struct {
	Index   string `json:"index"`
	Content string `json:"content"`
}

// Get all messages using a specific index.
func GetAllMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	vars := mux.Vars(request)
	index := vars["index"]

	messagesByIndex, err := messages.GetAllMessagesByIndex(nodeAddress, index)

	if err != nil {
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	} else {
		jsonInBytes, err := json.Marshal(messagesByIndex)

		if err != nil {
			jsonInString = "{\"error\": \"Unable to convert the messages struct into JSON format.\"}"
		} else {
			jsonInString = string(jsonInBytes) // TODO: Corrigir espaços em branco
		}
	}

	fmt.Fprint(writer, jsonInString)
}

// Create and submit a new message.
func CreateNewMessage(writer http.ResponseWriter, request *http.Request) {
	var message Message
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := "14265"
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	requestBody, _ := io.ReadAll(request.Body)

	json.Unmarshal(requestBody, &message)

	isMessageCreated := messages.SubmitMessage(nodeAddress, message.Index, message.Content, 15)

	if isMessageCreated {
		json.NewEncoder(writer).Encode(message)
	} else {
		log.Panic("Error during create a new message.")
	}
}
