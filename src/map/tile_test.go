package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tic-cat-dog-client/src/graphics"
)

func TestGenerateTiles(t *testing.T) {
	actual := GenerateTiles(320, 240, 3)

	expected := [3][3]Tile{
		{
			{UpperLeft: graphics.Point{X: 3, Y: 3}, BottomRight: graphics.Point{X: 106, Y: 80}},
			{UpperLeft: graphics.Point{X: 109, Y: 3}, BottomRight: graphics.Point{X: 212, Y: 80}},
			{UpperLeft: graphics.Point{X: 215, Y: 3}, BottomRight: graphics.Point{X: 318, Y: 80}},
		},

		{
			{UpperLeft: graphics.Point{X: 3, Y: 83}, BottomRight: graphics.Point{X: 106, Y: 160}},
			{UpperLeft: graphics.Point{X: 109, Y: 83}, BottomRight: graphics.Point{X: 212, Y: 160}},
			{UpperLeft: graphics.Point{X: 215, Y: 83}, BottomRight: graphics.Point{X: 318, Y: 160}},
		},

		{
			{UpperLeft: graphics.Point{X: 3, Y: 163}, BottomRight: graphics.Point{X: 106, Y: 240}},
			{UpperLeft: graphics.Point{X: 109, Y: 163}, BottomRight: graphics.Point{X: 212, Y: 240}},
			{UpperLeft: graphics.Point{X: 215, Y: 163}, BottomRight: graphics.Point{X: 318, Y: 240}},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestMarshalPoint(t *testing.T) {
	expected := graphics.Point{X: 12, Y: 15}
	actual := graphics.MarshalPoint(12, 15)

	assert.Equal(t, expected, actual)
}
