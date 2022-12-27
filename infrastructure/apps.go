package infrastructure

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"

	"github.com/afelinczak/app-get/domain"
)

const appFilePath = "/opt/app-get"
const appsFileName = "apps.json"

// EnsureIsAdmin will panic the app-get if it is not runned with root priviliges
func EnsureIsAdmin() bool {
	currentUser, _ := user.Current()
	if currentUser.Username != "root" {
		println("APP-GET requires root access. Switch to root account or use sudo app-get.")
		return false
	}
	return true
}

func getAppsFile() string {
	return appFilePath + "/" + appsFileName
}

// CreateInstalledAppListFile - creates empty apps.json file if not found
func CreateInstalledAppListFile() {
	var _, err = os.Stat(getAppsFile())

	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(appFilePath, 0744)
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
