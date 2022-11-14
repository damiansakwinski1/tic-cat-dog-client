package core

import "github.com/hajimehoshi/ebiten/v2"

type GameAdapter struct {
	Game *Game
}

func (gameAdapter *GameAdapter) Update() error {
	return gameAdapter.Game.Update()
}

func (gameAdapter *GameAdapter) Draw(screen *ebiten.Image) {
	gameAdapter.Game.Draw(screen)
}

func (gameAdapter *GameAdapter) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameAdapter.Game.Layout(outsideWidth, outsideHeight)
}
