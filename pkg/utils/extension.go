package utils

// IsJPEGish checks if the provided file extension is either "jpg" or "jpeg".
//
// It returns true if the extension matches "jpg" or "jpeg", otherwise false.
func IsJPEGish(extension string) bool {
	return extension == "jpg" || extension == "jpeg"
}

// IsPNG checks if the provided file extension is "png".
//
// It returns true if the extension matches "png", otherwise false.
func IsPNG(extension string) bool {
	return extension == "png"
}
