package infrastructure

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

const tempFilesDir string = "/tmp/"

func downloadFile(url string) (filePath string, err error) {

	var path string = tempFilesDir + path.Base(url)

	client := grab.NewClient()
	req, _ := grab.NewRequest(tempFilesDir, url)
	resp := client.Do(req)

	// start UI loop
	t := time.NewTicker(1000 * time.Millisecond)
	defer t.Stop()
	fmt.Println("")

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("\033[1A\033[K")
			fmt.Printf("  transferred %v kb / %v kb (%.2f%%)\n", resp.BytesComplete()/1024, resp.Size()/1024, 100*resp.Progress())
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}
	return path, nil
}
