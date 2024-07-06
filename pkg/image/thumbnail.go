package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/opa-oz/thumbnail/pkg/utils"
)

func ProcessImage(fpath string, width, height uint, newfname func(string) string) error {
	filename := filepath.Base(fpath)
	filenameParts := strings.Split(filename, ".")

	extension := filenameParts[1]

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

	m := resize.Thumbnail(width, height, img, resize.Lanczos3)

	err = utils.SaveImage(newfname(fpath), extension, &m)
	if err != nil {
		return err
	}

	return nil
}
