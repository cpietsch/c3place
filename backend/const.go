package main

import (
	"image"
)

const (
	imageDir    = "./static"
	imageWidth  = 1000
	imageHeight = 1000
)

var (
	upLeft   = image.Point{0, 0}
	lowRight = image.Point{imageWidth, imageHeight}
)
