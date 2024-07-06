package validators

import (
	"errors"
	"fmt"
	"regexp"
)

var validator, err = regexp.Compile("^[0-9]+x[0-9]+$")

func ValidateSize(size string) error {
	if err != nil {
		return err
	}

	if !validator.MatchString(size) {
		return errors.New(fmt.Sprintf("Incorrect `size` flag, please use `<width>x<height>` format"))
	}

	return nil
}
