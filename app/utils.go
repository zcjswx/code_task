package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func downloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFileToTmp(url string) (string, error) {
	timestamp := time.Now().UnixMilli()

	tmpPath := os.TempDir()
	fileName := fmt.Sprintf("graph-%s.xml", strconv.FormatInt(timestamp, 10))

	filePath := fmt.Sprintf("%s%s", tmpPath, fileName)
	err := downloadFile(url, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
