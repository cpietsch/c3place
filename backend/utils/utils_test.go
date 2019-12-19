package utils

import (
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
