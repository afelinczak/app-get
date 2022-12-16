package infrastructure

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/afelinczak/app-get/domain"
)

const APPS_FILE_PATH = "/home/{user}/.config/app-get"
const APPS_FILE_NAME = "apps.json"

func getAppsPath() string {
	currentUser, _ := user.Current()
	return strings.Replace(APPS_FILE_PATH, "{user}", currentUser.Username, 1)
}

func getAppsFile() string {
	return getAppsPath() + "/" + APPS_FILE_NAME
}

// CreateInstalledAppListFile - creates empty apps.json file if not found
func CreateInstalledAppListFile() {
	var _, err = os.Stat(getAppsFile())

	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(getAppsPath(), 0777)
		var newFileContent, errCreate = os.Create(getAppsFile())
		if errCreate != nil {
			log.Fatal(errCreate)
			os.Exit(101)
		}
		newFileContent.WriteString("[ ]")
		defer newFileContent.Close()
	}
}

// GetInstalledApps load list of apps from apps.json
func getInstalledApps() []domain.InstalledApp {
	var result = readFromFile(getAppsFile())
	var apps []domain.InstalledApp
	json.Unmarshal([]byte(result), &apps)
	return apps
}

// WriteInstalledApps overwrites list of installed apps
func writeInstalledApps(apps []domain.InstalledApp) {
	var json, _ = json.MarshalIndent(apps, "", "    ")
	writeToFile(getAppsFile(), string(json))
}

// AppRepository implementation
type AppRepository struct{}

// Get returns all installed applications
func (appRepo AppRepository) Get() []domain.InstalledApp {
	return getInstalledApps()
}

// Save overwrites list of apps
func (appRepo AppRepository) Save(apps []domain.InstalledApp) {
	writeInstalledApps(apps)
}
