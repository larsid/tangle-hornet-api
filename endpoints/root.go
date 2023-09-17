package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/allancapistrano/tangle-client-go/info"
)

// Shows information about Tangle Hornet Network
func Root(writer http.ResponseWriter, request *http.Request) {
	var jsonInString string
	nodeURL := "http://127.0.0.1:14265"

	// Network info
	nodeInfo, err := info.GetNodeInfo(nodeURL)

	if err != nil {
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	} else {
		json, err := json.Marshal(nodeInfo)

		if err != nil {
			jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		} else {
			jsonInString = string(json)
		}
	}

	fmt.Fprint(writer, jsonInString)
}
