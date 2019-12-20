package main

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/cpietsch/c3place/backend/pixel"
	"github.com/gin-gonic/gin"
)

/**
 * @api {post} /pixel
 * @apiDescription
 * create a new pixel
 * @apiExample Example usage:
 * curl -XPOST http://localhost:4000/ping
 * @apiSuccessExample {json} Success
{
	"status": "created",
	"pixel": {
		"x": 10,
		"y": 20,
		"r": 255,
		"g": 0,
		"b": 0
	}
}
*/
func handlerPixel(c *gin.Context) {
	body := pixel.PostPixel{}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post body"})
		return
	}

	err = pixel.ValidatePixel(body, imageWidth, imageHeight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("==> Write pixel pos (%v, %v) color (%v, %v, %v)\n", body.X, body.Y, body.R, body.G, body.B)
	data[body.X][body.Y] = color.RGBA{uint8(body.R), uint8(body.G), uint8(body.B), 0xff}

	newPixels = true

	c.JSON(http.StatusCreated, gin.H{"status": "created", "pixel": body})
}
