package main

import (
	"bytes"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerLatest(c *gin.Context) {
	if newPixels {
		img := buildImage()
		buf := new(bytes.Buffer)
		png.Encode(buf, img)
		imageCache = buf.Bytes()
	}
	c.Data(http.StatusOK, "image/png", imageCache)
}
