package main

import (
	"errors"
)

type PostPixel struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	X int   `json:"x"`
	Y int   `json:"y"`
}

func ValidatePixel(p PostPixel) error {
	if p.R < 0 || p.R > 255 {
		return errors.New("red not valid")
	}
	if p.G < 0 || p.G > 255 {
		return errors.New("green not valid")
	}
	if p.B < 0 || p.B > 255 {
		return errors.New("blue not valid")
	}
	if p.X < 0 || p.X > imageWidth {
		return errors.New("x not valid")
	}
	if p.Y < 0 || p.Y > imageHeight {
		return errors.New("y not valid")
	}
	return nil
}
