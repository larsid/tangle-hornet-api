package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/allancapistrano/tangle-client-go/info"
)

// Shows information about Tangle Hornet Network
func Root(writer http.ResponseWriter, request *http.Request) {
	nodeURL := "http://127.0.0.1:14265"

	// Network info
	nodeInfo := info.GetNodeInfo(nodeURL)

	json, err := json.Marshal(nodeInfo)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Fprint(writer, string(json))
}