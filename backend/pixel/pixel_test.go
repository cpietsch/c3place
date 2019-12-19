package pixel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePixel(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		pixel := PostPixel{R: 255, G: 255, B: 255, X: 0, Y: 0}
		err := ValidatePixel(pixel, 1000, 1000)
		assert.Nil(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		tests := []struct {
			desc string
			data PostPixel
			err  error
		}{
      // {
			// 	desc: "invalid r",
			// 	data: PostPixel{R: 500, G: 255, B: 255, X: 0, Y: 0},
			// 	err:  ErrRedNotValid,
			// },
      // {
			// 	desc: "invalid g",
			// 	data: PostPixel{R: 255, G: 500, B: 255, X: 0, Y: 0},
			// 	err:  ErrGreenNotValid,
			// },
      // {
			// 	desc: "invalid b",
			// 	data: PostPixel{R: 255, G: 255, B: 500, X: 0, Y: 0},
			// 	err:  ErrBlueNotValid,
			// },
      {
				desc: "invalid x",
				data: PostPixel{R: 255, G: 255, B: 255, X: -100, Y: 0},
				err:  ErrXNotValid,
			},
			{
				desc: "invalid y",
				data: PostPixel{R: 255, G: 255, B: 255, X: 0, Y: -100},
				err:  ErrYNotValid,
			},
		}

		for _, tt := range tests {
			t.Run(tt.desc, func(t *testing.T) {
				err := ValidatePixel(tt.data, 1000, 1000)
				assert.NotNil(t, err)
				assert.Equal(t, tt.err, err)
			})
		}
	})
}
