package cursor

import (
	"tic-cat-dog-client/src/graphics"
	board "tic-cat-dog-client/src/map"
)

func IsInTile(tile *board.Tile, pos graphics.Point) bool {
	return pos.X >= tile.UpperLeft.X &&
		pos.X <= tile.BottomRight.X &&
		pos.Y >= tile.UpperLeft.Y &&
		pos.Y <= tile.BottomRight.Y
}
