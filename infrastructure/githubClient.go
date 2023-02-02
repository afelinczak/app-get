package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
func GetInstallationFile(app *domain.App, source GithubVersion) (string, bool) {
	if app.AppType != domain.Deb {
		log.Fatal("Only deb files are supported")
		os.Exit(102)
	}

	if app.Source != domain.Github {
		log.Fatal("Only github source is supported")
		os.Exit(103)
	}

	appURLs := make(map[string]string)
	for i := 0; i < len(source.Assets); i++ {
		var parts []string = strings.Split(source.Assets[i].BrowserDownloadURL, "/")
		var fileName string = parts[len(parts)-1]

		if is64BitDebFile(fileName) == false {
			continue
		}
		if app.Name == "" {
			appURLs[fileName] = source.Assets[i].BrowserDownloadURL
		} else {
			match, _ := regexp.MatchString(app.Name+"[_-]{1}"+source.Name+"[_-]{1}", fileName)
			if match {
				appURLs[fileName] = source.Assets[i].BrowserDownloadURL
				break
			}
		}
	}

	keys := make([]string, 0, len(appURLs))
	for k := range appURLs {
		keys = append(keys, k)
	}

	appURL := ""

	if len(appURLs) == 0 {
		fmt.Println("  " + app.Name + " skip -> deb package not found")
		return "", false
	} else if len(appURLs) > 1 {
		name := askForFile(appURLs)
		app.Name = getShortAppName(name)
		fmt.Println("Adding " + app.Name + " to application list")
		appURL = appURLs[name]

	} else {
		for k, v := range appURLs {
			app.Name = getShortAppName(k)
			appURL = v
			break
		}
	}

	fmt.Println("Start downloading " + app.Name)
	var path, err = downloadFile(appURL)

	if err != nil {
		fmt.Println("Error while downloading deb file - skip")
		return "", false
	}
	return path, true
}

func is64BitDebFile(fileName string) bool {
	if strings.HasSuffix(fileName, "i686.deb") {
		return false
	}
	if strings.HasSuffix(fileName, "arm64.deb") {
		return false
	}
	if strings.HasSuffix(fileName, "armhf.deb") {
		return false
	}
	if strings.HasSuffix(fileName, "i686.deb") {
		return false
	}
	if strings.HasSuffix(fileName, ".deb") {
		return true
	}
	return false
}

func askForFile(files map[string]string) string {
	i := 1
	filesOptions := make([]string, 0)
	for fileName, _ := range files {
		fmt.Println(strconv.Itoa(i) + ". " + fileName)
		filesOptions = append(filesOptions, fileName)
		i++
	}

	_, err := fmt.Scanf("%d", &i)

	if err != nil || i > len(files)+1 || i < 1 {
		log.Fatal("Invalid option " + strconv.Itoa(i) + " - break")
		os.Exit(101)
	}

	fmt.Println(filesOptions[i-1])
	return filesOptions[i-1]
}

func getShortAppName(URL string) string {
	re := regexp.MustCompile("[_-]{1}v{0,1}[0-9]+")
	split := re.Split(URL, -1)
	return split[0]
}
