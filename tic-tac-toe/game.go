package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 420
	screenHeight = 600
	boardSize    = 3
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
	input      *Input

	boardX, boardY int
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBoard(boardSize),
	}
}

func (g *Game) Update() error {
	if g.isGameOver() {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.reset()
		}

		return nil
	}

	g.input.Update(g.boardX, g.boardY)
	g.board.Update(g.input)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	g.board.Draw(g.boardImage)

	drawBoardOpt := &ebiten.DrawImageOptions{}
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	boardWidth, boardHeight := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()

	drawBoardAtX := (screenWidth - boardWidth) / 2
	drawBoardAtY := (screenHeight - boardHeight) / 2

	if g.boardX == 0 {
		g.boardX = drawBoardAtX
	}
	if g.boardY == 0 {
		g.boardY = drawBoardAtY
	}

	drawBoardOpt.GeoM.Translate(float64(drawBoardAtX), float64(drawBoardAtY))
	screen.DrawImage(g.boardImage, drawBoardOpt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
}

func (g *Game) isGameOver() bool {
	return g.board.IsGameOver()
}

func (g *Game) reset() {
	g.board = NewBoard(boardSize)
	g.input = NewInput()
}
