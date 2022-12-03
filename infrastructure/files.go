package infrastructure

import (
	"io/ioutil"
	"log"
	"os"
)

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

func writeToFile(path string, content string) {
	var fileContent, errCreate = os.Create(path)
	if errCreate != nil {
		log.Fatal(errCreate)
		os.Exit(101)
	}
	fileContent.WriteString(content)
	defer fileContent.Close()
}
