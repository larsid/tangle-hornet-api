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

func GetAllMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	index := vars["index"]

	fmt.Fprintf(writer, "Retornando todas as mensagens do índice: "+index)
}

func CreateNewMessage(writer http.ResponseWriter, request *http.Request) {
	var message Message
	
	requestBody, _ := io.ReadAll(request.Body)

	json.Unmarshal(requestBody, &message)
	json.NewEncoder(writer).Encode(message)
}
