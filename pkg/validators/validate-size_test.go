package validators_test

import (
	"errors"
	"testing"

	"github.com/opa-oz/thumbnail/pkg/validators"
)

func TestValidateSize(t *testing.T) {
	tests := []struct {
		size     string
		expected error
	}{
		{"1920x1080", nil}, // Valid size format
		{"1024x768", nil},  // Another valid size format
		{"640x480", nil},   // Yet another valid size format
		{"1920x1080x720", errors.New("Incorrect `size` flag, please use `<width>x<height>` format")}, // Incorrect format
		{"1024x", errors.New("Incorrect `size` flag, please use `<width>x<height>` format")},         // Incorrect format
		{"x768", errors.New("Incorrect `size` flag, please use `<width>x<height>` format")},          // Incorrect format
	}

	for _, test := range tests {
		err := validators.ValidateSize(test.size)
		if err == nil && test.expected == nil {
			continue // Both are nil, test passed
		}
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err.Error() != test.expected.Error()) {
			t.Errorf("For size %s, expected error %v but got %v", test.size, test.expected, err)
		}
	}
}
