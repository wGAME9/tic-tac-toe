package tictactoe

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	blocks        [][]*Block
	size          int
	currentPlayer Player
	winner        Player
}

func NewBoard(size int) *Board {
	blocks := make([][]*Block, size)
	for i := range size {
		blocks[i] = make([]*Block, size)
		for j := range size {
			blocks[i][j] = NewBlock(i, j)
		}
	}

	return &Board{
		blocks:        blocks,
		size:          size,
		currentPlayer: playerX,
		winner:        playerUnknown,
	}
}

func (board *Board) nextPlayer() {
	if board.currentPlayer == playerX {
		board.currentPlayer = playerO
		return
	}

	board.currentPlayer = playerX
}

func (board *Board) Message() string {
	if board.HasWinner() {
		return fmt.Sprintf("Player %s wins!", board.Winner().String())
	}

	if board.IsDraw() {
		return "Draw!"
	}

	return fmt.Sprintf("Player %s's turn", board.currentPlayer.String())
}

func (board *Board) Update(input *Input) {
	if board.IsGameOver() {
		return
	}

	if !input.IsMouseJustPressed() {
		return
	}
	mouseX, mouseY := input.MousePosition()

	for i := range board.blocks {
		for j := range board.blocks[i] {
			block := board.blocks[i][j]
			cannnotSetThisBlock := !block.CanSetPlayer()
			if cannnotSetThisBlock {
				continue
			}

			isClickedOnThisBlock := block.BeingClicked(mouseX, mouseY)
			if !isClickedOnThisBlock {
				continue
			}

			block.SetPlayer(board.currentPlayer)

			if board.IsCurrentPlayerWin() {
				board.SetWinner(board.currentPlayer)
				return
			}

			board.nextPlayer()
			break
		}
	}
}

func (board *Board) Size() (int, int) {
	width := board.size*blockSize + (board.size+1)*blockMargin
	height := width

	return width, height
}

func (board *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)

	for i := range board.blocks {
		for j := range board.blocks[i] {
			board.blocks[i][j].Draw(boardImage)
		}
	}
}

func (board *Board) IsGameOver() bool {
	hasWinner := board.HasWinner()
	isDraw := board.IsDraw()
	return hasWinner || isDraw
}

func (board *Board) HasWinner() bool {
	return board.winner != playerUnknown
}

func (board *Board) Winner() Player {
	return board.currentPlayer
}

func (board *Board) SetWinner(player Player) {
	board.winner = player
}

func (board *Board) IsCurrentPlayerWin() bool {
	return board.isRowFilled(board.currentPlayer) ||
		board.isColFilled(board.currentPlayer) ||
		board.isDiagonalFilled(board.currentPlayer)
}

func (board *Board) IsDraw() bool {
	for i := range board.blocks {
		for j := range board.blocks[i] {
			if board.blocks[i][j].Player() == playerUnknown {
				return false
			}
		}
	}

	return true
}

func (board *Board) isRowFilled(player Player) bool {
	for i := range board.blocks {
		filled := true
		for j := range board.blocks[i] {
			if board.blocks[i][j].Player() != player {
				filled = false
				break
			}
		}

		if filled {
			return true
		}
	}

	return false
}

func (board *Board) isColFilled(player Player) bool {
	for j := range board.blocks {
		filled := true
		for i := range board.blocks[j] {
			if board.blocks[i][j].Player() != player {
				filled = false
				break
			}
		}

		if filled {
			return true
		}
	}

	return false
}

func (board *Board) isDiagonalFilled(player Player) bool {
	filled := true
	for i := range board.blocks {
		if board.blocks[i][i].Player() != player {
			filled = false
			break
		}
	}

	if filled {
		return true
	}

	filled = true
	for i := range board.blocks {
		if board.blocks[i][board.size-1-i].Player() != player {
			filled = false
			break
		}
	}

	return filled
}
