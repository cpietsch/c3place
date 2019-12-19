package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://yourbasic.org/golang/create-image/
func handlerLatest(c *gin.Context) {
	c.Data(http.StatusOK, "image/png", imageBuffer.Bytes())
}
