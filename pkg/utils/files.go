package utils

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// IsFileExists checks if the file specified by the filepath exists.
//
// It returns true if the file exists, otherwise false. It handles errors
// using os.Stat to determine existence, returning false if the file does
// not exist or if an error occurs.
func IsFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// SaveImage saves the image pointed to by m to the specified filepath.
//
// It saves the image in the format specified by the extension parameter.
// Currently supported formats include JPEG ("jpg" or "jpeg") and PNG ("png").
// Returns an error if there is an issue creating the file or encoding the image.
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
