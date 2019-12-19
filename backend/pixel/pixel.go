package pixel

import (
	"errors"
)

var (
	ErrRedNotValid   = errors.New("red not valid")
	ErrGreenNotValid = errors.New("green not valid")
	ErrBlueNotValid  = errors.New("blue not valid")
	ErrXNotValid     = errors.New("x not valid")
	ErrYNotValid     = errors.New("y not valid")
)

type PostPixel struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	X int   `json:"x"`
	Y int   `json:"y"`
}

func ValidatePixel(p PostPixel, imageWidth, imageHeight int) error {
	if p.R < 0 || p.R > 255 {
		return ErrRedNotValid
	}
	if p.G < 0 || p.G > 255 {
		return ErrGreenNotValid
	}
	if p.B < 0 || p.B > 255 {
		return ErrBlueNotValid
	}
	if p.X < 0 || p.X > imageWidth {
		return ErrXNotValid
	}
	if p.Y < 0 || p.Y > imageHeight {
		return ErrYNotValid
	}
	return nil
}
