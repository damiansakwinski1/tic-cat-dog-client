package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
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
