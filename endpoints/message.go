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
	Index   string      `json:"index"`
	Content interface{} `json:"content"`
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
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		jsonInBytes, err := json.Marshal(messagesByIndex)

		if err != nil {
			jsonInString = "{\"error\": \"Unable to convert the messages struct into JSON format.\"}"
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(jsonInBytes)
		}
	}

	fmt.Fprint(writer, jsonInString)
}

// Get a message by given message ID.
func GetMessageByMessageId(writer http.ResponseWriter, request *http.Request) {
	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	vars := mux.Vars(request)
	messageId := vars["messageID"]

	message, err := messages.GetMessageFormattedByMessageID(nodeAddress, messageId)
	if err != nil {
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		jsonInBytes, err := json.Marshal(message)

		if err != nil {
			jsonInString = "{\"error\": \"Unable to convert the messages struct into JSON format.\"}"
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(jsonInBytes)
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

	// Serializar o objeto para formato JSON
	jsonContent, err := json.Marshal(message.Content)
	if err != nil {
		errorMessage := fmt.Sprintf("Error serializing the object: %s", err.Error())
		jsonInString := fmt.Sprintf("{\"error\": \"%s\"}", errorMessage)

		log.Println(errorMessage)
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		// Converter o JSON em uma string
		contentString := string(jsonContent)

		isMessageCreated := messages.SubmitMessage(nodeAddress, message.Index, contentString, 15)

		if isMessageCreated {
			json.NewEncoder(writer).Encode(message)
		} else {
			errorMessage := "Error during create a new message."
			jsonInString := fmt.Sprintf("{\"error\": \"%s\"}", errorMessage)

			log.Println(errorMessage)
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		}
	}
}
