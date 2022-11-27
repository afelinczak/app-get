package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const TEMP_FILES_DIR string = "/tmp/"

func downloadFile(url string) (filePath string, err error) {

	var path string = TEMP_FILES_DIR + path.Base(url)
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}
