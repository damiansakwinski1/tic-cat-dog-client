package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"os"
	"tic-cat-dog-client/src/cursor"
	"tic-cat-dog-client/src/graphics"
	board "tic-cat-dog-client/src/map"
)

type Image interface {
	Size() (width, height int)
	Clear()
	Fill(clr color.Color)
	DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions)
	DrawTriangles(vertices []ebiten.Vertex, indices []uint16, img *ebiten.Image, options *ebiten.DrawTrianglesOptions)
	DrawTrianglesShader(vertices []ebiten.Vertex, indices []uint16, shader *ebiten.Shader, options *ebiten.DrawTrianglesShaderOptions)
	DrawRectShader(width, height int, shader *ebiten.Shader, options *ebiten.DrawRectShaderOptions)
	SubImage(r image.Rectangle) image.Image
	Bounds() image.Rectangle
	ColorModel() color.Model
	ReadPixels(pixels []byte)
	At(x, y int) color.Color
	RGBA64At(x, y int) color.RGBA64
	Set(x, y int, clr color.Color)
	Dispose()
	WritePixels(pixels []byte)
	ReplacePixels(pixels []byte)
}

type MouseButtonJustPressedChecker func(button ebiten.MouseButton) bool
type GetCursorPosition func() (x int, y int)
type DrawRect func(screen *ebiten.Image, x float64, y float64, width float64, height float64, clr color.Color)

type State struct {
	MouseButtonJustPressed MouseButtonJustPressedChecker
	GetCursorPosition      GetCursorPosition
	DrawRect               DrawRect
}

type Drafter interface {
	DrawRect(img Image, x float64, y float64, width float64, height float64, clr color.Color)
}

type DrafterImpl struct{}

func (d *DrafterImpl) DrawRect(img Image, x float64, y float64, width float64, height float64, clr color.Color) {
	ebitenutil.DrawRect(img.(*ebiten.Image), x, y, width, height, clr)
}

type Game struct {
	Board        board.Board
	State        State
	Drafter      Drafter
	DrawCat      bool
	ScreenWidth  uint
	ScreenHeight uint
}

func (g *Game) UpdateBoard() {
	if !g.State.MouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}

	for rowKey, row := range g.Board {
		for cellKey := range row {
			if !cursor.IsInTile(&g.Board[rowKey][cellKey], graphics.MarshalPoint(g.State.GetCursorPosition())) {
				continue
			}

			g.Board[rowKey][cellKey].HasCharacter = true
		}
	}
}

func (g *Game) Update() error {
	g.UpdateBoard()

	return nil
}

/*
*
TODO: Refactor so that:
 1. It is possible to mock opening an image.
 2. The method is shorter.
*/
func (g *Game) Draw(screen Image) {
	basePath := os.Getenv("ASSETS_ABS_PATH")
	catImg, _, err := ebitenutil.NewImageFromFile(basePath + "/cat.png")
	dogImg, _, err := ebitenutil.NewImageFromFile(basePath + "/dog.png")
	catImgOptions := ebiten.DrawImageOptions{}
	//catImgOptions.GeoM.Translate(-30, -30)
	catImgOptions.GeoM.Scale(0.15, 0.15)
	dogImgOptions := ebiten.DrawImageOptions{}
	dogImgOptions.GeoM.Scale(0.15, 0.15)
	dogImgOptions.GeoM.Translate(100, 0)

	if err != nil {
		HandleError(err)
	}

	var barThickness float64 = 3
	screen.Fill(colornames.White)
	if g.DrawCat {
		screen.DrawImage(catImg, &catImgOptions)
	}
	screen.DrawImage(dogImg, &dogImgOptions)

	//Board prototype
	g.Drafter.DrawRect(screen, 0, 0, 320, barThickness, colornames.Black)
	g.Drafter.DrawRect(screen, 0, 0, 320, barThickness, colornames.Black)
	g.Drafter.DrawRect(screen, 0, float64(g.ScreenHeight)/3, 320, barThickness, colornames.Black)
	g.Drafter.DrawRect(screen, 0, float64(g.ScreenHeight)*(float64(2)/float64(3)), 320, barThickness, colornames.Black)
	g.Drafter.DrawRect(screen, 0, 0, barThickness, float64(g.ScreenHeight), colornames.Black)
	g.Drafter.DrawRect(screen, float64(g.ScreenWidth)-barThickness, 0, barThickness, float64(g.ScreenHeight), colornames.Black)
	g.Drafter.DrawRect(screen, float64(g.ScreenWidth)/3, 0, barThickness, float64(g.ScreenHeight), colornames.Black)
	g.Drafter.DrawRect(screen, float64(g.ScreenWidth)*(float64(2)/float64(3)), 0, barThickness, float64(g.ScreenHeight), colornames.Black)
	g.Drafter.DrawRect(screen, 0, float64(g.ScreenHeight)-barThickness, float64(g.ScreenWidth), barThickness, colornames.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.ScreenWidth), int(g.ScreenHeight)
}
