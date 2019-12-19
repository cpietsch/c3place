package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{2, 2}})
	img.Set(0, 0, color.RGBA{255, 0, 0, 0xff})
	img.Set(1, 0, color.RGBA{0, 255, 0, 0xff})
	img.Set(0, 1, color.RGBA{0, 0, 255, 0xff})
	img.Set(1, 1, color.RGBA{0, 0, 0, 0xff})
	f, _ := os.Create("3.png")
	png.Encode(f, img)
}
