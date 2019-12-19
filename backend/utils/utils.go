package utils

import (
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func FilenameWithoutExtension(p string) string {
	return strings.TrimSuffix(p, path.Ext(p))
}

func GetLatestImageFilename(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	newestFile := ""
	newestTime := 0
	for _, f := range files {
		if f.Mode().IsRegular() {
			filename := f.Name()
			if filepath.Ext(filename) == ".png" {
				currTime, err := strconv.Atoi(FilenameWithoutExtension(filename))
				if err != nil {
					return "", err
				} else if currTime > newestTime {
					newestTime = currTime
					newestFile = filename
				}
			}
		}
	}
	return newestFile, nil
}

func LoadPngToColorArray(filename string, w, h int) ([]color.RGBA, error) {
	data := make([]color.RGBA, 0)

	existingImageFile, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer existingImageFile.Close()
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		return data, err
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := loadedImage.At(x, y)
			r, g, b, a := c.RGBA()
			data = append(data, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	return data, nil
}
