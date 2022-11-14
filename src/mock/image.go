package mock

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/mock"
	"image"
	"image/color"
)

type Image struct {
	ebiten.Image
	mock.Mock
}

func (mi *Image) Size() (width, height int) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) Clear() {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) Fill(clr color.Color) {
	mi.Called(clr)
}

func (mi *Image) DrawTriangles(vertices []ebiten.Vertex, indices []uint16, img *ebiten.Image, options *ebiten.DrawTrianglesOptions) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) DrawTrianglesShader(vertices []ebiten.Vertex, indices []uint16, shader *ebiten.Shader, options *ebiten.DrawTrianglesShaderOptions) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) DrawRectShader(width, height int, shader *ebiten.Shader, options *ebiten.DrawRectShaderOptions) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) SubImage(r image.Rectangle) image.Image {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) Bounds() image.Rectangle {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) ColorModel() color.Model {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) ReadPixels(pixels []byte) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) At(x, y int) color.Color {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) RGBA64At(x, y int) color.RGBA64 {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) Set(x, y int, clr color.Color) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) Dispose() {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) WritePixels(pixels []byte) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) ReplacePixels(pixels []byte) {
	//TODO implement me
	panic("implement me")
}

func (mi *Image) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {
	mi.Called(img, options)
}
