package install

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Create a temporary file and download the installer to it.
func DownloadFile(installerName string, url string, filename string) (string, error) {

	fmt.Println("Downloading " + installerName + " installer from: " + url)

	// Create the file
	tmpFile, err := os.CreateTemp("", filename)
	if err != nil {
		return tmpFile.Name(), err
	}
	defer tmpFile.Close()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New("error creating request")
	}
	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("error downloading " + installerName + " installer")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errors.New("error retrieving " + installerName + " installer")
	}

	// Writer the body to file
	_, err = io.Copy(tmpFile, res.Body)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}
