package board

import "tic-cat-dog-client/src/graphics"

type Character string

type Tile struct {
	UpperLeft    graphics.Point
	BottomRight  graphics.Point
	HasCharacter bool
	Character    Character
}

func (t *Tile) GetCenter() graphics.Point {
	panic("Implement me")
}
