package image

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/opa-oz/thumbnail/pkg/utils"
)

var validExtensions = []string{"jpg", "jpeg", "png"}

func ProcessImage(fpath string, width, height uint, newfname func(string) string) error {
	if !utils.IsFileExists(fpath) {
		return errors.New(fmt.Sprintf("File %s doesn't exist", fpath))
	}

	filename := filepath.Base(fpath)
	filenameParts := strings.Split(filename, ".")

	if len(filenameParts) != 2 {
		return errors.New(fmt.Sprintf("File %s doesn't have extension", filename))
	}

	valid := false
	extension := filenameParts[1]

	for _, ext := range validExtensions {
		if extension == ext {
			valid = true
		}
	}

	if !valid {
		return errors.New(fmt.Sprintf("File extension %s is not supported", extension))
	}

	file, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image

	if utils.IsJPEGish(extension) {
		img, err = jpeg.Decode(file)
		if err != nil {
			return err
		}
	} else if utils.IsPNG(extension) {
		img, err = png.Decode(file)
		if err != nil {
			return err
		}
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)

	err = utils.SaveImage(newfname(fpath), extension, &m)
	if err != nil {
		return err
	}

	return nil
}
