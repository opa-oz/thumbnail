package validators

import (
	"errors"
	"fmt"

	"github.com/opa-oz/thumbnail/pkg/utils"
)

var validExtensions = []string{"jpg", "jpeg", "png"}

func ExistsOrError(fpath string) error {
	if !utils.IsFileExists(fpath) {
		return errors.New(fmt.Sprintf("File %s doesn't exist", fpath))
	}

	return nil
}

func ExtensionOrError(parts *[]string, filename string) error {
	if len(*parts) != 2 {
		return errors.New(fmt.Sprintf("File %s doesn't have extension", filename))
	}

	return nil
}

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
