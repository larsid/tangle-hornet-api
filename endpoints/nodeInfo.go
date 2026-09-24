package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/larsid/tangle-client-go/info"
	"github.com/larsid/tangle-hornet-api/config"
)

const CONFIG_FILE_NAME = "tangle-hornet.conf"

// Shows information about Tangle Hornet Network
func GetNodeInfo(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), 30*time.Second)
	defer cancel()

	log.Printf("[API-HANDLER] [INFO] GetNodeInfo: requisição recebida")

	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	// Network info
	nodeInfo, err := info.GetNodeInfo(ctx, nodeAddress)

	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetNodeInfo: retornando HTTP 500 error=%v", err)
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		json, err := json.Marshal(nodeInfo)

		if err != nil {
			log.Printf("[API-HANDLER] [ERROR] GetNodeInfo: falha ao serializar JSON error=%v", err)
			jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(json)
			log.Printf("[API-HANDLER] [INFO] GetNodeInfo: resposta HTTP 200 healthy=%v pruning_index=%d payload_bytes=%d",
				nodeInfo.IsHealthy, nodeInfo.Milestone.PruningIndex, len(jsonInString))
		}
	}

	fmt.Fprint(writer, jsonInString)
}

// Shows all information about Tangle Hornet Network
func GetAllNodeInfo(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), 30*time.Second)
	defer cancel()

	log.Printf("[API-HANDLER] [INFO] GetAllNodeInfo: requisição recebida")

	var jsonInString string
	nodeURL := config.GetNodeUrl(CONFIG_FILE_NAME, true)
	nodePort := config.GetNodePort(CONFIG_FILE_NAME, true)
	nodeAddress := fmt.Sprintf("http://%s:%s", nodeURL, nodePort)

	// All network info
	allNodeInfo, err := info.GetAllNodeInfo(ctx, nodeAddress)

	if err != nil {
		log.Printf("[API-HANDLER] [ERROR] GetAllNodeInfo: retornando HTTP 500 error=%v", err)
		jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		http.Error(writer, jsonInString, http.StatusInternalServerError)
	} else {
		json, err := json.Marshal(allNodeInfo)

		if err != nil {
			log.Printf("[API-HANDLER] [ERROR] GetAllNodeInfo: falha ao serializar JSON error=%v", err)
			jsonInString = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
			http.Error(writer, jsonInString, http.StatusInternalServerError)
		} else {
			jsonInString = string(json)
			log.Printf("[API-HANDLER] [INFO] GetAllNodeInfo: resposta HTTP 200 healthy=%v pruning_index=%d payload_bytes=%d",
				allNodeInfo.IsHealthy, allNodeInfo.Milestone.PruningIndex, len(jsonInString))
		}
	}

	fmt.Fprint(writer, jsonInString)
}
