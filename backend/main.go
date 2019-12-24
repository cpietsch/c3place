package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"time"

	"github.com/cpietsch/c3place/backend/config"
	"github.com/cpietsch/c3place/backend/utils"
)

var (
	cfg = config.Config{}

	// in-mem data
	data           [][]color.RGBA
	dataGroundplan [][]color.RGBA
	newPixels      bool
	imageCache     []byte
)

func main() {
	log.Printf("c3place v%s\n\n", version)

	setupData()

	// concurrent persist images
	go persistImages(imageDir)

	// start the server
	r := setupRouter()
	r.Run(":" + cfg.Port)
}

func setupData() {
	// load the groundplan
	var err error
	dataGroundplan, err = utils.LoadPngToColorArray("groundplan.png", imageWidth, imageHeight)
	if err != nil {
		panic(err)
	}

	// load the last image and add data to the data array
	latestImage, err := utils.GetLatestImageFilename(imageDir)
	if err != nil {
		panic(err)
	}
	data, err = utils.LoadPngToColorArray(path.Join(imageDir, latestImage), imageWidth, imageHeight)
	if err != nil {
		panic(err)
	}
}

func buildImage() image.Image {
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			// draw the groundplan
			if dataGroundplan[x][y].R == 255 &&
				dataGroundplan[x][y].G == 255 &&
				dataGroundplan[x][y].B == 255 {
				img.Set(x, y, colorGroundplan)
			} else {
				// draw the pixels
				img.Set(x, y, data[x][y])
			}
		}
	}

	return img
}

func persistImages(dir string) {
	if newPixels {
		img := buildImage()
		buf := new(bytes.Buffer)
		png.Encode(buf, img)
		imageCache = buf.Bytes()

		now := time.Now()
		filename := path.Join(dir, strconv.Itoa(int(now.Unix()))+".png")
		err := ioutil.WriteFile(filename, imageCache, 0755)
		if err != nil {
			log.Println("Error write png", err)
		}
		log.Println("==> PNG file written", filename)

		newPixels = false
	}

	time.Sleep(5 * time.Second)
	persistImages(dir)
}
