package cursor_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tic-cat-dog-client/src/cursor"
	"tic-cat-dog-client/src/graphics"
	board "tic-cat-dog-client/src/map"
)

type TestParameters struct {
	Expected       bool
	CursorPosition graphics.Point
	Tile           board.Tile
}

func TestIsInTile(t *testing.T) {
	testParameters := []TestParameters{
		{
			Expected:       true,
			Tile:           board.Tile{UpperLeft: graphics.Point{X: 3, Y: 3}, BottomRight: graphics.Point{X: 106, Y: 80}},
			CursorPosition: graphics.Point{X: 4, Y: 7},
		},
		{
			Expected:       false,
			Tile:           board.Tile{UpperLeft: graphics.Point{X: 109, Y: 3}, BottomRight: graphics.Point{X: 212, Y: 80}},
			CursorPosition: graphics.Point{X: 4, Y: 7},
		},
		{
			Expected:       true,
			Tile:           board.Tile{UpperLeft: graphics.Point{X: 109, Y: 3}, BottomRight: graphics.Point{X: 212, Y: 80}},
			CursorPosition: graphics.Point{X: 109, Y: 3},
		},
	}

	for _, testParameter := range testParameters {
		actual := cursor.IsInTile(&testParameter.Tile, testParameter.CursorPosition)
		assert.Equal(t, testParameter.Expected, actual)
	}
}
