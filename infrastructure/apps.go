package infrastructure

import (
	"errors"
	"log"
	"os"
)

const APP_FILE_PATH = "apps.json"

func CreateInstalledAppListFile() {
	var _, err = os.Stat(APP_FILE_PATH)

	if errors.Is(err, os.ErrNotExist) {
		var newFileContent, errCreate = os.Create(APP_FILE_PATH)
		if errCreate != nil {
			log.Fatal(errCreate)
			os.Exit(101)
		}
		newFileContent.WriteString("[ ]")
		defer newFileContent.Close()
	}
}
