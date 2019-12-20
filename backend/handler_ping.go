package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /ping
 * @apiDescription
 * test if the api is alive
 * @apiExample Example usage:
 * curl -XGET http://localhost:4000/ping
 * @apiSuccessExample {json} Success
{
	"ping": "pong"
}
*/
func handlerPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
