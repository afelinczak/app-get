package domain

import "fmt"

type AppType string

const (
	Deb      AppType = "deb"
	AppImage AppType = "appImage"
)

type AppSource string

const (
	Github AppSource = "github"
)

type App struct {
	Name      string    `json:"name"`
	AppUrl    string    `json:"app_url"`
	AppType   AppType   `json:"app_type"`
	Source    AppSource `json:"source"`
	SourceUrl string    `json:"source_url"`
}

// AddNewApp - saves to disk updated list of installed apps
func AddNewApp(newApp App, version string, appsRepo IAppsRepository) {
	var apps = appsRepo.Get()
	var newAppWithVersion = InstalledApp{App: newApp, Version: version}
	var newList = append(apps, newAppWithVersion)
	appsRepo.Save(newList)
}

// RemoveApp - saves to disk updated list of installed apps
func RemoveApp(repoName string, appsRepo IAppsRepository) {
	var apps = appsRepo.Get()

	for i := 0; i < len(apps); i++ {
		if apps[i].App.SourceUrl == repoName {
			fmt.Println("Removing repository " + repoName)
			newList := append(apps[:i], apps[i+1:]...)
			appsRepo.Save(newList)
			fmt.Println("You can now remove app using sudo apt-get remove " + apps[i].App.Name)
			return
		}
	}
	fmt.Println("Repository " + repoName + " not found on list of apps.")
}

// UpdateAppVersion - updates version of the app and saves to disk
func UpdateAppVersion(updatedApp App, version string, appsRepo IAppsRepository) {
	var apps = appsRepo.Get()
	for i := 0; i < len(apps); i++ {
		if apps[i].App.Name == updatedApp.Name {
			apps[i].Version = version
			break
		}
	}
	appsRepo.Save(apps)
}
