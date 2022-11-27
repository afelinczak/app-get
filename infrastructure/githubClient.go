package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/afelinczak/app-get/domain"
)

const GITHUB_API_URL string = "https://api.github.com/"

type GitHubAsset struct {
	Name               string `json:"name"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type GithubVersion struct {
	Name   string        `json:"name"`
	Assets []GitHubAsset `json:"assets"`
}

func GetLatestVersion(app domain.App) GithubVersion {
	var url = GITHUB_API_URL + "repos/" + app.SourceUrl + "/releases/latest"
	var resp, err = http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var appVersion GithubVersion
	json.Unmarshal([]byte(body), &appVersion)

	if strings.HasPrefix(appVersion.Name, "v") {
		appVersion.Name = strings.Replace(appVersion.Name, "v", "", 1)
	}

	return appVersion
}

func GetInstallationFile(app domain.App, source GithubVersion) string {
	if app.AppType != domain.Deb {
		log.Fatal("Only deb files are supported")
		os.Exit(102)
	}

	if app.Source != domain.Github {
		log.Fatal("Only github source is supported")
		os.Exit(103)
	}

	var appUrl string
	for i := 0; i < len(source.Assets); i++ {
		var parts []string = strings.Split(source.Assets[i].BrowserDownloadUrl, "/")
		var fileName string = parts[len(parts)-1]

		if strings.HasSuffix(fileName, "amd64.deb") {
			var regVer string = strings.Replace(source.Name, ".", "\\.", 3)
			match, _ := regexp.MatchString(app.Name+"[_-]{1}"+regVer+"[_-]{1}"+"amd64.deb", fileName)
			if match {
				appUrl = source.Assets[i].BrowserDownloadUrl
				break
			}
		}
	}

	if appUrl == "" {
		log.Fatal("amd64 deb file not found")
		os.Exit(404)
	}

	var path, err = downloadFile(appUrl)

	if err != nil {
		log.Fatalln(err)
		os.Exit(105)
	}
	return path
}