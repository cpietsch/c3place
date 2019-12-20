package main

import (
	"log"
	"os"
	"time"

	"github.com/cpietsch/c3place/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"github.com/zsais/go-gin-prometheus"
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

	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)

	// initialize the routes
	router.GET("/ping", handlerPing)
	router.GET("/latest", handlerLatest)
	router.POST("/pixel", handlerPixel)

	// initialize the static server
	router.StaticFS("/timetravel", gin.Dir("./static", true))

	return router
}
