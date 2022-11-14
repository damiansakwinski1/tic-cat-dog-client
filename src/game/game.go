package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"tic-cat-dog-client/src/core"
	board "tic-cat-dog-client/src/map"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

func Init() *core.Game {
	return &core.Game{
		Board:        board.GenerateTiles(ScreenWidth, ScreenHeight, 3),
		ScreenWidth:  ScreenWidth,
		ScreenHeight: ScreenHeight,
		Drafter:      &core.DrafterImpl{},
		State:        core.State{MouseButtonJustPressed: inpututil.IsMouseButtonJustPressed, GetCursorPosition: ebiten.CursorPosition},
	}
}
