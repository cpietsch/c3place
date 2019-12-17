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
	port        string
	redisHost   string
	redisPort   string
	rateLimiter bool

	imageWidth  = 1000
	imageHeight = 1000
	upLeft      = image.Point{0, 0}
	lowRight    = image.Point{imageWidth, imageHeight}

	data []PostPixel
)

func setupRouter() *gin.Engine {
	// get env vars
	port = os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisPort = os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	rateLimiterEnv := os.Getenv("RATELIMITER")
	if rateLimiterEnv == "true" {
		rateLimiter = true
	} else {
		rateLimiter = false
	}
	log.Printf("HOST        : %s\n", port)
	log.Printf("REDIS_HOST  : %s\n", redisHost)
	log.Printf("REDIS_PORT  : %s\n", redisPort)
	log.Printf("RATELIMITER : %v\n", rateLimiter)

	// Create a redis client.
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Println("REDIS CLIENT:", client)
	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_gin",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// setup the router
	router := gin.Default()

	// Create a new middleware with the limiter instance.
	if rateLimiter {
		rate := limiter.Rate{
			Period: 1 * time.Second,
			Limit:  100,
		}
		middleware := mgin.NewMiddleware(limiter.New(store, rate))
		router.ForwardedByClientIP = true
		router.Use(middleware)
	}

	// initialize the routes
	router.GET("/ping", handlerPing)
	router.GET("/", handlerIndex)
	router.POST("/pixel", handlerPixel)

	return router
}

func main() {
	// start the server
	r := setupRouter()
	r.Run(":" + port)
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
