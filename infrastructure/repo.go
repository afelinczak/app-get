package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/afelinczak/app-get/domain"
)

const REPO_PATH string = "repo.json"
const APPS_PATH string = "apps.json"

func GetRepo() []domain.App {
	var result = readFromFile(REPO_PATH)
	var apps []domain.App
	json.Unmarshal([]byte(result), &apps)
	return apps
}

func GetInstalledApps() []domain.InstalledApp {

	var result = readFromFile(APPS_PATH)
	var apps []domain.InstalledApp
	json.Unmarshal([]byte(result), &apps)
	return apps
}

func readFromFile(path string) []byte {

	fileContent, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
		os.Exit(100)
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)
	return byteResult
}
