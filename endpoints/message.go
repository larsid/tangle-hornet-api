package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllMessagesByIndex(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	index:= vars["index"]

	fmt.Fprintf(writer, "Retornando todas as mensagens do índice: " + index)
}