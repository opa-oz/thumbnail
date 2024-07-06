package validators

import (
	"errors"
	"fmt"
)

var validExtensions = []string{"jpg", "jpeg", "png"}

// FileExistsFunc is an interface that abstracts the IsFileExists function.
type FileExistsFunc func(fpath string) bool

// ExistsOrError checks if the file specified by fpath exists.
//
// It returns an error if the file does not exist.
func ExistsOrError(fpath string, isFileExists FileExistsFunc) error {
	if !isFileExists(fpath) {
		return errors.New(fmt.Sprintf("File %s doesn't exist", fpath))
	}

	return nil
}

// ExtensionOrError checks if the filename has a valid extension.
//
// It expects parts to contain two elements: the filename and its extension.
// Returns an error if the filename doesn't have a valid extension format.
func ExtensionOrError(parts *[]string, filename string) error {
	if len(*parts) != 2 {
		return errors.New(fmt.Sprintf("File %s doesn't have extension", filename))
	}

	return nil
}

// SupportedOrError checks if the provided extension is among the supported ones.
//
// It compares the extension parameter against a predefined list of valid extensions
// (jpg, jpeg, png). Returns an error if the extension is not supported.
func SupportedOrError(extension string) error {
	valid := false

	for _, ext := range validExtensions {
		if extension == ext {
			valid = true
		}
	}

	if !valid {
		return errors.New(fmt.Sprintf("File extension %s is not supported", extension))
	}

	return nil
}
