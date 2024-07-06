package utils_test

import (
	"testing"

	"github.com/opa-oz/thumbnail/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsJPEGish(t *testing.T) {
	tests := []struct {
		extension string
		expected  bool
	}{
		{"jpg", true},
		{"jpeg", true},
		{"JPG", false},      // Case-sensitive check
		{"png", false},      // Different extension
		{"", false},         // Empty extension
		{"jpeg2000", false}, // Different format
	}

	for _, test := range tests {
		result := utils.IsJPEGish(test.extension)
		assert.Equal(t, result, test.expected)
	}
}

func TestIsPNG(t *testing.T) {
	tests := []struct {
		extension string
		expected  bool
	}{
		{"png", true},
		{"PNG", false},       // Case-sensitive check
		{"jpg", false},       // Different extension
		{"", false},          // Empty extension
		{"jpeg", false},      // Different format
		{"image/png", false}, // Different MIME type
	}

	for _, test := range tests {
		result := utils.IsPNG(test.extension)
		assert.Equal(t, result, test.expected)
	}
}
