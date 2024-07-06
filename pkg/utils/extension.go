package utils

func IsJPEGish(extension string) bool {
	return extension == "jpg" || extension == "jpeg"
}

func IsPNG(extension string) bool {
	return extension == "png"
}
