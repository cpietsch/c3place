package main

import (
	"image"
	"image/color"
)

const (
	imageDir    = "./static"
	imageWidth  = 1000
	imageHeight = 1000
)

var (
	upLeft          = image.Point{0, 0}
	lowRight        = image.Point{imageWidth, imageHeight}
	colorGroundplan = color.RGBA{0, 255, 0, 0xff}
)
