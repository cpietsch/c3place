package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
