package board

import "tic-cat-dog-client/src/graphics"

func GenerateTiles(screenWidth int, screenHeight int, tileGap int) [3][3]Tile {
	var result [3][3]Tile
	tileWidth := screenWidth / 3
	tileHeight := screenHeight / 3

	for i, upperLeftY := 0, tileGap; i < 3; i, upperLeftY = i+1, upperLeftY+tileHeight {
		for j, upperLeftX := 0, tileGap; j < 3; j, upperLeftX = j+1, upperLeftX+tileWidth {
			result[i][j] = Tile{
				UpperLeft:   graphics.Point{X: upperLeftX, Y: upperLeftY},
				BottomRight: graphics.Point{X: upperLeftX + tileWidth - tileGap, Y: upperLeftY + tileHeight - tileGap},
			}
		}
	}

	return result
}
