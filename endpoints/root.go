package endpoints

import (
	"fmt"
	"net/http"
)

// TODO: Colocar para exibir as informações da rede Tangle
func Root(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home page")
}