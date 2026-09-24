package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/larsid/tangle-client-go/messages"
	"github.com/larsid/tangle-hornet-api/config"
)

type Message struct {
	Index string      `json:"index"`
	Data  interface{} `json:"data"`
}

const readMessagesTimeout = 2 * time.Minute
const writeMessageTimeout = 30 * time.Second

// Get all messages using a specific index.
// Optional query parameter: limit (positive integer) returns only the most recent N messages.
func GetAllMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), readMessagesTimeout)
	defer cancel()

	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	vars := mux.Vars(request)
	index := vars["index"]
	limitParam := request.URL.Query().Get("limit")

	log.Printf("[API-HANDLER] [INFO] GetAllMessagesByIndex: requisição recebida index=%s limit=%s", index, limitParam)

	var messagesByIndex []messages.Message
	var err error

	if limitParam != "" {
		limit, parseErr := strconv.Atoi(limitParam)
		if parseErr != nil || limit <= 0 {
			log.Printf("[API-HANDLER] [ERROR] GetAllMessagesByIndex: parâmetro limit inválido index=%s limit=%s", index, limitParam)
			http.Error(
				writer,
				`{"error": "Invalid limit parameter. Must be a positive integer."}`,
				http.StatusBadRequest,
			)
			return
		}
		messagesByIndex, err = messages.GetMessagesByIndexWithLimit(ctx, nodeAddress, index, limit)
	} else {
		messagesByIndex, err = messages.GetAllMessagesByIndex(ctx, nodeAddress, index)
	}

	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetAllMessagesByIndex: retornando HTTP 500 index=%s error=%v", index, err)
		http.Error(
			writer,
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	jsonInBytes, err := json.Marshal(messagesByIndex)
	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetAllMessagesByIndex: falha ao serializar JSON index=%s error=%v", index, err)
		http.Error(
			writer,
			`{"error": "Unable to convert the messages struct into JSON format."}`,
			http.StatusInternalServerError,
		)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(writer, string(jsonInBytes))
	log.Printf("[API-HANDLER] [INFO] GetAllMessagesByIndex: resposta HTTP 200 index=%s mensagens=%d payload_bytes=%d", index, len(messagesByIndex), len(jsonInBytes))
}

// Get a limited amount of messages created in the last hour, available on the
// node by a given index.
func GetLastHourMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), readMessagesTimeout)
	defer cancel()

	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	vars := mux.Vars(request)
	index := vars["index"]
	maxMessagesInString := vars["maxMessages"]

	log.Printf("[API-HANDLER] [INFO] GetLastHourMessagesByIndex: requisição recebida index=%s maxMessages=%s", index, maxMessagesInString)

	maxMessages, err := strconv.Atoi(maxMessagesInString)
	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetLastHourMessagesByIndex: maxMessages inválido index=%s maxMessages=%s", index, maxMessagesInString)
		jsonInString = "{\"error\": \"Invalid maximum number of messages.\"}"
		http.Error(writer, jsonInString, http.StatusBadRequest)
		return
	}

	lastHourMessagesByIndex, err := messages.GetLastHourMessagesByIndex(
		ctx,
		nodeAddress,
		index,
		maxMessages,
	)

	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetLastHourMessagesByIndex: retornando HTTP 500 index=%s error=%v", index, err)
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		jsonInBytes, err := json.Marshal(lastHourMessagesByIndex)

		if err != nil {
			log.Printf("[API-HANDLER] [ERROR] GetLastHourMessagesByIndex: falha ao serializar JSON index=%s error=%v", index, err)
			jsonInString = "{\"error\": \"Unable to convert the messages struct into JSON format.\"}"
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(jsonInBytes)
			log.Printf("[API-HANDLER] [INFO] GetLastHourMessagesByIndex: resposta HTTP 200 index=%s mensagens=%d payload_bytes=%d", index, len(lastHourMessagesByIndex), len(jsonInBytes))
		}
	}

	fmt.Fprint(writer, jsonInString)
}

// Get a message by given message ID.
func GetMessageByMessageId(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), 30*time.Second)
	defer cancel()

	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	vars := mux.Vars(request)
	messageId := vars["messageID"]

	log.Printf("[API-HANDLER] [INFO] GetMessageByMessageId: requisição recebida messageID=%s", messageId)

	message, err := messages.GetMessageFormattedByMessageID(ctx, nodeAddress, messageId)
	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetMessageByMessageId: retornando HTTP 500 messageID=%s error=%v", messageId, err)
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		jsonInBytes, err := json.Marshal(message)

		if err != nil {
			log.Printf("[API-HANDLER] [ERROR] GetMessageByMessageId: falha ao serializar JSON messageID=%s error=%v", messageId, err)
			jsonInString = "{\"error\": \"Unable to convert the messages struct into JSON format.\"}"
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(jsonInBytes)
			log.Printf("[API-HANDLER] [INFO] GetMessageByMessageId: resposta HTTP 200 messageID=%s payload_bytes=%d", messageId, len(jsonInBytes))
		}
	}

	fmt.Fprint(writer, jsonInString)
}

// Create and submit a new message.
func CreateNewMessage(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), writeMessageTimeout)
	defer cancel()

	var message Message
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := "14265"
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	requestBody, _ := io.ReadAll(request.Body)

	json.Unmarshal(requestBody, &message)

	log.Printf("[API-HANDLER] [INFO] CreateNewMessage: requisição recebida index=%s payload_bytes=%d", message.Index, len(requestBody))

	// Serializar o objeto para formato JSON
	jsonContent, err := json.Marshal(message.Data)
	if err != nil {
		errorMessage := fmt.Sprintf("Error serializing the object: %s", err.Error())
		jsonInString := fmt.Sprintf("{\"error\": \"%s\"}", errorMessage)

		log.Printf("[API-HANDLER] [ERROR] CreateNewMessage: retornando HTTP 500 index=%s error=%s", message.Index, errorMessage)
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		// Converter o JSON em uma string
		contentString := string(jsonContent)

		isMessageCreated := messages.SubmitMessage(ctx, nodeAddress, message.Index, contentString, 15)

		if isMessageCreated {
			json.NewEncoder(writer).Encode(message)
			log.Printf("[API-HANDLER] [INFO] CreateNewMessage: resposta HTTP 200 index=%s mensagem criada com sucesso", message.Index)
		} else {
			errorMessage := "Error during create a new message."
			jsonInString := fmt.Sprintf("{\"error\": \"%s\"}", errorMessage)

			log.Printf("[API-HANDLER] [ERROR] CreateNewMessage: retornando HTTP 500 index=%s error=%s", message.Index, errorMessage)
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		}
	}
}
