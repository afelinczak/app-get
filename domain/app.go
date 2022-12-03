package domain

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
func AddNewApp(newApp App, version string, getApps func() []InstalledApp, save func([]InstalledApp)) {
	var apps = getApps()
	var newAppWithVersion = InstalledApp{App: newApp, Version: version}
	var newList = append(apps, newAppWithVersion)
	save(newList)
}
