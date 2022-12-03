package infrastructure

import (
	"encoding/json"

	"github.com/afelinczak/app-get/domain"
)

const REPO_PATH string = "repo.json"
const APPS_PATH string = "apps.json"

// GetRepo load list of available apps from repo.json
func GetRepo() []domain.App {
	var result = readFromFile(REPO_PATH)
	var apps []domain.App
	json.Unmarshal([]byte(result), &apps)
	return apps
}
