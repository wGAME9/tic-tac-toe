package tictactoe

import "github.com/hajimehoshi/ebiten/v2"

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type Input struct {
	mouseState     mouseState
	mousePositionX int
	mousePositionY int
}

func NewInput() *Input {
	return &Input{
		mouseState: mouseStateNone,
	}
}

func (input *Input) Update(boardX, boardY int) {
	switch input.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			input.mouseState = mouseStatePressing
			input.mousePositionX = x
			input.mousePositionY = y

			input.normalize(boardX, boardY)
		}

	case mouseStatePressing:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			input.mouseState = mouseStateSettled
		}

	case mouseStateSettled:
		input.Reset()
	}
}

// normalize normalizes the mouse position to the actual board position
// since the board size is smaller than the game window size, so the user
// may clicked outside the board area.
func (input *Input) normalize(boardX, boardY int) {
	input.mousePositionX -= boardX
	input.mousePositionY -= boardY
}

func (input *Input) Reset() {
	input.mouseState = mouseStateNone
	input.mousePositionX = 0
	input.mousePositionY = 0
}

func (input *Input) IsMouseJustPressed() bool {
	return input.mouseState == mouseStatePressing
}

func (input *Input) MousePosition() (int, int) {
	return input.mousePositionX, input.mousePositionY
}
