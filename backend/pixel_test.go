package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePixel(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		pixel := PostPixel{R: 255, G: 255, B: 255, X: 0, Y: 0}
		err := ValidatePixel(pixel)
		assert.Nil(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		pixel := PostPixel{R: 255, G: 255, B: 255, X: -100, Y: 0}
		err := ValidatePixel(pixel)
		assert.NotNil(t, err)
		assert.Equal(t, "x not valid", err.Error())
	})
}
