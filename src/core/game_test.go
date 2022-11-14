package core_test

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"golang.org/x/image/colornames"
	"testing"
	"tic-cat-dog-client/src/core"
	"tic-cat-dog-client/src/graphics"
	board "tic-cat-dog-client/src/map"
	"tic-cat-dog-client/src/mock"
)

// TODO: Finish this tests so that it covers 100% of game.Draw
func TestDraw(t *testing.T) {
	state := core.State{
		MouseButtonJustPressed: func(button ebiten.MouseButton) bool {
			return false
		},
	}

	testAdapter := new(mock.Drafter)
	testAdapter.On("DrawRect", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).Return(nil)
	game := core.Game{State: state, Drafter: testAdapter}
	mockScreen := new(mock.Image)

	mockScreen.On("DrawImage", mock2.Anything, mock2.Anything).Return(nil)
	mockScreen.On("Fill", colornames.White).Return(nil)
	game.Draw(mockScreen)
	mockScreen.AssertCalled(t, "Fill", colornames.White)
	testAdapter.AssertCalled(t, "DrawRect", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything)
}

func TestUpdateBoard(t *testing.T) {
	type TestParameters struct {
		CursorPosition         graphics.Point
		ExpectedBoard          board.Board
		MouseButtonJustPressed bool
	}

	var expectedBoards []board.Board

	for i := 0; i < 3; i++ {
		expectedBoards = append(expectedBoards, board.GenerateTiles(320, 240, 3))
	}

	expectedBoards[0][0][0].HasCharacter = true
	expectedBoards[1][2][2].HasCharacter = true

	testParameters := []TestParameters{
		{CursorPosition: graphics.Point{X: 5, Y: 11}, ExpectedBoard: expectedBoards[0], MouseButtonJustPressed: true},
		{CursorPosition: graphics.Point{X: 218, Y: 165}, ExpectedBoard: expectedBoards[1], MouseButtonJustPressed: true},
		{CursorPosition: graphics.Point{X: 150, Y: 83}, ExpectedBoard: expectedBoards[2], MouseButtonJustPressed: false},
	}

	for _, testParameter := range testParameters {
		state := core.State{
			MouseButtonJustPressed: func(button ebiten.MouseButton) bool {
				return testParameter.MouseButtonJustPressed
			},
			GetCursorPosition: func() (x int, y int) {
				return testParameter.CursorPosition.X, testParameter.CursorPosition.Y
			},
		}

		game := core.Game{
			State: state,
			Board: board.GenerateTiles(320, 240, 3),
		}

		game.UpdateBoard()
		assert.Equal(t, testParameter.ExpectedBoard, game.Board)
	}
}
