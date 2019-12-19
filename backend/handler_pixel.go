package main

import (
	"fmt"
	"net/http"

	"github.com/cpietsch/c3place/backend/pixel"
	"github.com/gin-gonic/gin"
)

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

	fmt.Println("==> Write pixel to data", body)
	data = append(data, body)

	newPixels = true

	c.JSON(http.StatusCreated, gin.H{"status": "created", "pixel": body})
}
