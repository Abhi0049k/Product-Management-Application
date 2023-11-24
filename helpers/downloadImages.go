package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// To Download all the images present in product_images

func DownloadImages(imageURLs []string) ([]string, []string, error) {
	var localFilePaths []string

	var compressedFilePaths []string

	var err error

	for _, imageURL := range imageURLs {
		fileName := fmt.Sprintf("image_%d%s", time.Now().UnixNano(), filepath.Ext(imageURL))
		destination := filepath.Join("./images", fileName)
		destination = filepath.ToSlash(destination)
		localFilePaths = append(localFilePaths, destination)
		// Downloading the image in images folder
		err = DownloadImage(imageURL, destination)
		if err != nil {
			return nil, nil, err
		}

		compressedFileName := fmt.Sprintf("compressed_%d%s", time.Now().UnixNano(), filepath.Ext(imageURL))
		compressedDestination := filepath.Join("./compressed-images", compressedFileName)
		compressedDestination = filepath.ToSlash(compressedDestination)
		err = os.MkdirAll(filepath.Dir(compressedDestination), os.ModePerm)
		if err != nil {
			return nil, nil, err
		}

		// Compressing and downloading the image in compressed folder
		err = CompressImage(destination, compressedDestination, 100, 100)
		if err != nil {
			return nil, nil, err
		}
		compressedFilePaths = append(compressedFilePaths, compressedDestination)
	}

	return localFilePaths, compressedFilePaths, nil
}
