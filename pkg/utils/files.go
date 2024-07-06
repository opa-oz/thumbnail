package utils

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func IsFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func SaveImage(filepath string, extension string, m *image.Image) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	if IsJPEGish(extension) {
		jpeg.Encode(out, *m, nil)
	} else if IsPNG(extension) {
		png.Encode(out, *m)
	}

	return nil
}
