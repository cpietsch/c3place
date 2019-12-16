package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
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
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10,
	}

	// Create a redis client.
	option, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	client := redis.NewClient(option)

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_gin_example",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create a new middleware with the limiter instance.
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(middleware)

	router.GET("/ping", handlerPing)
	router.GET("/", handlerIndex)
	router.POST("/pixel", handlerPixel)

	return router
}

func main() {
	r := setupRouter()
	r.Run(port)
}

func handlerPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// https://yourbasic.org/golang/create-image/
func handlerIndex(c *gin.Context) {
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i < len(data); i++ {
		img.Set(data[i].X, data[i].Y, color.RGBA{data[i].R, data[i].G, data[i].B, 0xff})
	}

	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

func handlerPixel(c *gin.Context) {
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
}
