package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Get the port on which the API will run.
func GetApiPort(fileName string, debug bool) string {
	apiPort := "3000"
	foundSettingsFile := true
	directory := "/etc/"

	filePath := filepath.Join(directory, fileName)
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
