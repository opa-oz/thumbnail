package validators

import (
	"errors"
	"fmt"
	"regexp"
)

// validator is a regular expression compiled to match `<width>x<height>` format.
var validator, err = regexp.Compile("^[0-9]+x[0-9]+$")

// ValidateSize validates the format of the size string.
//
// It checks if the size string matches the `<width>x<height>` format using a regular expression.
// Returns an error if the size string is empty, doesn't match the expected format, or if there was an error compiling the regular expression.
func ValidateSize(size string) error {
	if err != nil {
		return err
	}

	if !validator.MatchString(size) {
		return errors.New(fmt.Sprintf("Incorrect `size` flag, please use `<width>x<height>` format"))
	}

	return nil
}
