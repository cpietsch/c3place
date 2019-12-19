package utils

import (
	"io/ioutil"
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
