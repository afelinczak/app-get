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
