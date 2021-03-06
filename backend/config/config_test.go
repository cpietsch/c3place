package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInit(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		cfg := ConfigInit()
		assert.Equal(t, "4000", cfg.Port)
		assert.Equal(t, "localhost", cfg.RedisHost)
		assert.Equal(t, "6379", cfg.RedisPort)
		assert.Equal(t, false, cfg.RateLimiter)
	})

	t.Run("env values", func(t *testing.T) {
		os.Setenv("PORT", "7357")
		os.Setenv("REDIS_HOST", "10.0.0.1")
		os.Setenv("REDIS_PORT", "9876")
		os.Setenv("RATELIMITER", "true")

		cfg := ConfigInit()
		assert.Equal(t, "7357", cfg.Port)
		assert.Equal(t, "10.0.0.1", cfg.RedisHost)
		assert.Equal(t, "9876", cfg.RedisPort)
		assert.Equal(t, true, cfg.RateLimiter)

		cfg.Log()
	})

}
