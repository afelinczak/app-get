package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/afelinczak/app-get/domain"
)

const githubAPIURL string = "https://api.github.com/"

type gitHubAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

// GithubVersion contains latest version with list of assets
type GithubVersion struct {
	Name   string        `json:"tag_name"`
	Assets []gitHubAsset `json:"assets"`
}

// GetLatestVersion loads latest stable version of an app from github
func GetLatestVersion(app domain.App) GithubVersion {
	var url = githubAPIURL + "repos/" + app.SourceUrl + "/releases/latest"
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

// GetInstallationFile will download amd64 deb file to disk. Returns path and succes
func GetInstallationFile(app domain.App, source GithubVersion) (string, bool) {
	if app.AppType != domain.Deb {
		log.Fatal("Only deb files are supported")
		os.Exit(102)
	}

	if app.Source != domain.Github {
		log.Fatal("Only github source is supported")
		os.Exit(103)
	}

	var appURL string
	for i := 0; i < len(source.Assets); i++ {
		var parts []string = strings.Split(source.Assets[i].BrowserDownloadURL, "/")
		var fileName string = parts[len(parts)-1]

		if strings.HasSuffix(fileName, "amd64.deb") {
			var regVer string = strings.Replace(source.Name, ".", "\\.", 3)
			match, _ := regexp.MatchString(app.Name+"[_-]{1}"+regVer+"[_-]{1}"+"amd64.deb", fileName)
			if match {
				appURL = source.Assets[i].BrowserDownloadURL
				break
			}
		}
	}

	if appURL == "" {
		fmt.Println("  " + app.Name + " skip -> deb package not found")
		return "", false
	}

	var path, err = downloadFile(appURL)

	if err != nil {
		fmt.Println("Error while downloading deb file - skip")
		return "", false
	}
	return path, true
}
