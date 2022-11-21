package domain

type AppType string

const (
	deb      AppType = "deb"
	appImage AppType = "appImage"
)

type AppSource string

const (
	github AppSource = "github"
)

type App struct {
	name      string
	url       string
	appUrl    string
	appType   AppType
	source    AppSource
	sourceUrl string
}
