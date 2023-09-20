package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const CONFIG_DIRECTORY = "/etc/"

// Get the port on which the API will run.
func GetApiPort(fileName string, debug bool) string {
	apiPort := "3000"
	foundSettingsFile := true

	filePath := filepath.Join(CONFIG_DIRECTORY, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		if debug {
			fmt.Printf(
				"Couldn't open the '%s' file! Check the path or file name.\n",
				fileName,
			)
		}
		foundSettingsFile = false
	}

	defer file.Close()

	if foundSettingsFile {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			// Ignore the comments.
			if strings.Contains(scanner.Text(), "#") {
				continue
			}

			// Found API port setting.
			if strings.Contains(scanner.Text(), "apiPort") {
				apiPort = strings.Split(scanner.Text(), " = ")[1]
			}
		}
	}

	return apiPort
}

// Get the node URL used by the Tangle Hornet node.
func GetNodeUrl(fileName string, debug bool) string {
	nodeUrl := "127.0.0.1"
	foundSettingsFile := true

	filePath := filepath.Join(CONFIG_DIRECTORY, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		if debug {
			fmt.Printf(
				"Couldn't open the '%s' file! Check the path or file name.\n",
				fileName,
			)
		}
		foundSettingsFile = false
	}

	defer file.Close()

	if foundSettingsFile {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			// Ignore the comments.
			if strings.Contains(scanner.Text(), "#") {
				continue
			}

			// Found node URL setting.
			if strings.Contains(scanner.Text(), "nodeUrl") {
				nodeUrl = strings.Split(scanner.Text(), " = ")[1]
			}
		}
	}

	return nodeUrl
}
