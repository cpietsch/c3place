package utils

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilenameWithoutExtension(t *testing.T) {
	result := FilenameWithoutExtension("test.png")
	assert.Equal(t, "test", result)
}

func TestGetLatestImageFilename(t *testing.T) {
	t.Run("without error", func(t *testing.T) {
		result, err := GetLatestImageFilename("./fixtures")
		assert.Nil(t, err)
		assert.Equal(t, "3.png", result)
	})

	t.Run("with error", func(t *testing.T) {
		result, err := GetLatestImageFilename("./wrong")
		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})
}

func TestLoadPngToColorArray(t *testing.T) {
	t.Run("without error", func(t *testing.T) {
		data, err := LoadPngToColorArray("./fixtures/3.png", 2, 2)
		assert.Nil(t, err)
		assert.Len(t, data, 4)
		assert.Equal(t, color.RGBA{255, 0, 0, 0xff}, data[0])
		assert.Equal(t, color.RGBA{0, 0, 255, 0xff}, data[1])
		assert.Equal(t, color.RGBA{0, 255, 0, 0xff}, data[2])
		assert.Equal(t, color.RGBA{0, 0, 0, 0xff}, data[3])
	})

	t.Run("with error", func(t *testing.T) {
		_, err := LoadPngToColorArray("./wrong/3.png", 2, 2)
		assert.NotNil(t, err)
	})
}
