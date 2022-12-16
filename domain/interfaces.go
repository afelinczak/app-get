package domain

// IAppsRepository allows to access apps
type IAppsRepository interface {
	Get() []InstalledApp
	Save([]InstalledApp)
}
