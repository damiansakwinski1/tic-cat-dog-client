package mock

import (
	"github.com/stretchr/testify/mock"
	"image/color"
	"tic-cat-dog-client/src/graphics"
)

type Drafter struct {
	mock.Mock
}

func (d *Drafter) DrawRect(img graphics.Image, x float64, y float64, width float64, height float64, clr color.Color) {
	d.Called(img, x, y, width, height, clr)
}
