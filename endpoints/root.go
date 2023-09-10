package endpoints

import (
	"fmt"
	"net/http"
)

func Root(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Home page")
}