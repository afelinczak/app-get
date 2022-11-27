package domain

type InstalledApp struct {
	App     App    `json:"app"`
	Version string `json:"version"`
}
