package main

import (
	"bytes"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerLatest(c *gin.Context) {
	if newPixels {
		img, imgGroundplan := buildImage()

		buf := new(bytes.Buffer)
		png.Encode(buf, img)
		cacheImage = buf.Bytes()

		bufGroundplan := new(bytes.Buffer)
		png.Encode(bufGroundplan, imgGroundplan)
		cacheImageGroundplate = bufGroundplan.Bytes()
	}
	c.Data(http.StatusOK, "image/png", cacheImageGroundplate)
}
