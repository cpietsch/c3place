package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/cpietsch/c3place/backend/config"
	"github.com/cpietsch/c3place/backend/pixel"
	"github.com/cpietsch/c3place/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

var (
	version = "0.1.1"

	cfg = config.Config{}

	imageWidth  = 1000
	imageHeight = 1000
	upLeft      = image.Point{0, 0}
	lowRight    = image.Point{imageWidth, imageHeight}

	data        []pixel.PostPixel
	newPixels   bool
	imageBuffer = new(bytes.Buffer)
)

func setupRouter() *gin.Engine {
	cfg = config.ConfigInit()
	cfg.Log()

	// setup the router
	router := gin.Default()
	router.Use(cors.Default())
	// Create a new middleware with the limiter instance.
	if cfg.RateLimiter {
		// Create a redis client.
		client := redis.NewClient(&redis.Options{
			Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
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
	router.GET("/latest", handlerLatest)
	router.POST("/pixel", handlerPixel)

	// initialize the static server
	router.StaticFS("/timetravel", gin.Dir("./static", true))

	return router
}

func main() {
	log.Printf("c3place v%s\n\n", version)

	latestImage, err := utils.GetLatestImageFilename("./static")
	if err != nil {
		panic(err)
	}

	loadPngToData(latestImage)

	go persistImages("./static")

	// start the server
	r := setupRouter()
	r.Run(":" + cfg.Port)
}

func loadPngToData(filename string) {
	existingImageFile, err := os.Open("./static/" + filename)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	// fmt.Println(loadedImage)
	// bounds := loadedImage.Bounds()
	// w, h := bounds.Max.X, bounds.Max.Y
	w, h := imageWidth, imageHeight
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			color := loadedImage.At(x, y)
			r, g, b, _ := color.RGBA()
			pixel := pixel.PostPixel{R: uint8(r), G: uint8(g), B: uint8(b), X: x, Y: y}
			data = append(data, pixel)
		}
	}
	newPixels = true
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
		png.Encode(imageBuffer, img)
		log.Println("==> write png file", filename)
		newPixels = false
	}

	time.Sleep(5 * time.Second)
	persistImages(dir)
}
