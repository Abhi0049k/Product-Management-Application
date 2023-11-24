package helpers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// To Download the image and store it in images folder

func DownloadImage(url, destination string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = os.MkdirAll(filepath.Dir(destination), os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}
	return nil
}
