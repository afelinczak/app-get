package infrastructure

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/afelinczak/app-get/domain"
)

const APP_FILE_PATH = "apps.json"

// CreateInstalledAppListFile - creates empty apps.json file if not found
func CreateInstalledAppListFile() {
	var _, err = os.Stat(APP_FILE_PATH)

	if errors.Is(err, os.ErrNotExist) {
		var newFileContent, errCreate = os.Create(APP_FILE_PATH)
		if errCreate != nil {
			log.Fatal(errCreate)
			os.Exit(101)
		}
		newFileContent.WriteString("[ ]")
		defer newFileContent.Close()
	}
}

// GetInstalledApps load list of apps from apps.json
func GetInstalledApps() []domain.InstalledApp {
	var result = readFromFile(APPS_PATH)
	var apps []domain.InstalledApp
	json.Unmarshal([]byte(result), &apps)
	return apps
}

// WriteInstalledApps overwrites list of installed apps
func WriteInstalledApps(apps []domain.InstalledApp) {
	var json, _ = json.MarshalIndent(apps, "", "    ")
	writeToFile(APPS_PATH, string(json))
}
