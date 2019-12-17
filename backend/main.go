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
	"path"
	"strconv"
	"time"

	"github.com/cpietsch/c3place/backend/pixel"
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

	data      []pixel.PostPixel
	newPixels bool
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

	// setup the router
	router := gin.Default()

	// Create a new middleware with the limiter instance.
	if rateLimiter {
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

	// initialize the static server
	router.StaticFS("/timetravel", gin.Dir("./static", true))

	return router
}

func main() {
	go persistImages("./static")

	// start the server
	r := setupRouter()
	r.Run(":" + port)
}

func handlerPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// https://yourbasic.org/golang/create-image/
func handlerIndex(c *gin.Context) {
	img := buildImage()
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

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

func buildImage() image.Image {
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < len(data); i++ {
		img.Set(data[i].X, data[i].Y, color.RGBA{data[i].R, data[i].G, data[i].B, 0xff})
	}
	return img
}

func persistImages(dir string) {
	if newPixels {
		img := buildImage()
		now := time.Now()
		filename := path.Join(dir, strconv.Itoa(int(now.Unix()))+".png")
		f, _ := os.Create(filename)
		png.Encode(f, img)
		log.Println("==> write png file", filename)
		newPixels = false
	}

	time.Sleep(5 * time.Second)
	persistImages(dir)
}
