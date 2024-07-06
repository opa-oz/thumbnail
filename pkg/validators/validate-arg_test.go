package validators_test

import (
	"fmt"
	"testing"

	"github.com/opa-oz/thumbnail/pkg/validators"
)

// Mock implementation of IsFileExists for testing ExistsOrError function.
func mockIsFileExists(fpath string) bool {
	// Mocking file existence based on fpath
	return fpath == "existing_file.txt"
}

func TestExistsOrError(t *testing.T) {
	tests := []struct {
		fpath    string
		expected error
	}{
		{"existing_file.txt", nil}, // Existing file
		{"non_existing_file.txt", fmt.Errorf("File non_existing_file.txt doesn't exist")}, // Non-existing file
	}

	for _, test := range tests {
		err := validators.ExistsOrError(test.fpath, mockIsFileExists)
		if err == nil && test.expected == nil {
			continue // Both are nil, test passed
		}
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err.Error() != test.expected.Error()) {
			t.Errorf("For fpath %s, expected error %v but got %v", test.fpath, test.expected, err)
		}
	}
}

func TestExtensionOrError(t *testing.T) {
	tests := []struct {
		parts    []string
		filename string
		expected error
	}{
		{[]string{"filename", "jpg"}, "filename.jpg", nil},                                     // Valid extension
		{[]string{"filename"}, "filename", fmt.Errorf("File filename doesn't have extension")}, // Missing extension
	}

	for _, test := range tests {
		err := validators.ExtensionOrError(&test.parts, test.filename)
		if err == nil && test.expected == nil {
			continue // Both are nil, test passed
		}
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err.Error() != test.expected.Error()) {
			t.Errorf("For filename %s, expected error %v but got %v", test.filename, test.expected, err)
		}
	}
}

func TestSupportedOrError(t *testing.T) {
	tests := []struct {
		extension string
		expected  error
	}{
		{"jpg", nil},  // Supported extension
		{"jpeg", nil}, // Supported extension
		{"png", nil},  // Supported extension
		{"txt", fmt.Errorf("File extension txt is not supported")}, // Unsupported extension
	}

	for _, test := range tests {
		err := validators.SupportedOrError(test.extension)
		if err == nil && test.expected == nil {
			continue // Both are nil, test passed
		}
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err.Error() != test.expected.Error()) {
			t.Errorf("For extension %s, expected error %v but got %v", test.extension, test.expected, err)
		}
	}
}
