package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	RedisHost   string
	RedisPort   string
	RateLimiter bool
}

// ConfigInit get the configuration over env vars
func ConfigInit() Config {
	c := Config{}

	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		c.Port = "4000"
	}
	c.RedisHost = os.Getenv("REDIS_HOST")
	if c.RedisHost == "" {
		c.RedisHost = "localhost"
	}
	c.RedisPort = os.Getenv("REDIS_PORT")
	if c.RedisPort == "" {
		c.RedisPort = "6379"
	}
	rateLimiterEnv := os.Getenv("RATELIMITER")
	if rateLimiterEnv == "true" {
		c.RateLimiter = true
	} else {
		c.RateLimiter = false
	}

	return c
}

// Log the config values
func (c *Config) Log() {
	log.Printf("HOST        : %s\n", c.Port)
	log.Printf("REDIS_HOST  : %s\n", c.RedisHost)
	log.Printf("REDIS_PORT  : %s\n", c.RedisPort)
	log.Printf("RATELIMITER : %v\n", c.RateLimiter)
}
