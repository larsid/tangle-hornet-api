package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larsid/tangle-client-go/info"
	"github.com/allancapistrano/tangle-hornet-api/config"
)

const CONFIG_FILE_NAME = "tangle-hornet.conf"

// Shows information about Tangle Hornet Network
func GetNodeInfo(writer http.ResponseWriter, request *http.Request) {
	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	// Network info
	nodeInfo, err := info.GetNodeInfo(nodeAddress)

	if err != nil {
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		json, err := json.Marshal(nodeInfo)

		if err != nil {
			jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(json)
		}
	}

	fmt.Fprint(writer, jsonInString)
}

// Shows all information about Tangle Hornet Network
func GetAllNodeInfo(writer http.ResponseWriter, request *http.Request) {
	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	// All network info
	allNodeInfo, err := info.GetAllNodeInfo(nodeAddress)

	if err != nil {
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		json, err := json.Marshal(allNodeInfo)

		if err != nil {
			jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(json)
		}
	}

	fmt.Fprint(writer, jsonInString)
}