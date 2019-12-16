package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	port = ":4000"

	imageWidth  = 1000
	imageHeight = 1000
	upLeft      = image.Point{0, 0}
	lowRight    = image.Point{imageWidth, imageHeight}

	data []PostPixel
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// https://yourbasic.org/golang/create-image/
	r.GET("/", func(c *gin.Context) {
		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

		for i := 0; i < len(data); i++ {
			img.Set(data[i].X, data[i].Y, color.RGBA{data[i].R, data[i].G, data[i].B, 0xff})
		}

		buf := new(bytes.Buffer)
		png.Encode(buf, img)
		c.Data(http.StatusOK, "image/png", buf.Bytes())
	})

	r.POST("/pixel", func(c *gin.Context) {
		body := PostPixel{}

		err := c.Bind(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post body"})
			return
		}

		err = ValidatePixel(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("==> Write pixel to data", body)
		data = append(data, body)

		c.JSON(http.StatusCreated, gin.H{"status": "created", "pixel": body})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(port)
}
