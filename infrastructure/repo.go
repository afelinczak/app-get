package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/afelinczak/app-get/domain"
)

const REPO_PATH string = "repo.json"

func GetRepo() {

	//pwd, _ := os.Getwd()
	fileContent, err := os.Open(REPO_PATH)

	if err != nil {
		log.Fatal(err)
		os.Exit(100)
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var repo []domain.App
	json.Unmarshal([]byte(byteResult), &repo)

	fmt.Println(repo)
}
