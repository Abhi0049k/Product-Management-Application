package helpers

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// To Compress and download the image

func CompressImage(source, destination string, maxWidth, maxHeight uint) error {
	file, err := os.Open(source)

	if err != nil {
		return err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	resizedImg := resize.Resize(maxWidth, maxHeight, img, resize.Lanczos3)

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, resizedImg, nil)
	if err != nil {
		return err
	}

	return nil
}
