package main

import (
	"log"

	tictactoe "game/tic-tac-toe/tic-tac-toe"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 420
	screenHeight = 600
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tic Tac Toe")

	game := tictactoe.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
