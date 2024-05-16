package tictactoe

import (
	"fmt"

	"github.com/wGAME9/animation/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	lineSpeed   = 10
	lineWidth   = 5
	linePadding = 10
)

type Board struct {
	blocks        [][]*Block
	size          int
	currentPlayer Player
	winner        Player

	winningLine                animation.Animation
	isRowWinning               bool
	winningOnRow               int
	isColWinning               bool
	winningOnCol               int
	isDiagonalWinning          bool
	winningOnDiagonalDirection diagonalDirection
}

func NewBoard(size int) *Board {
	blocks := make([][]*Block, size)
	for rowNum := range size {
		blocks[rowNum] = make([]*Block, size)
		for colNum := range size {
			blocks[rowNum][colNum] = NewBlock(rowNum, colNum)
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
	if board.winningLine != nil {
		board.winningLine.Update()
	}

	if board.IsGameOver() {
		board.setWinningLine()
		return
	}

	if !input.IsMouseJustPressed() {
		return
	}
	mouseX, mouseY := input.MousePosition()

	for row := range board.blocks {
		for col := range board.blocks[row] {
			block := board.blocks[row][col]
			cannotSetThisBlock := !block.CanSetPlayer()
			if cannotSetThisBlock {
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

	for row := range board.blocks {
		for col := range board.blocks[row] {
			board.blocks[row][col].Draw(boardImage)
		}
	}

	if board.winningLine != nil {
		winningLineImg := board.winningLine.Image()
		boardImage.DrawImage(winningLineImg, &ebiten.DrawImageOptions{})
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
	isRowFilled, rowNum := board.isRowFilled(board.currentPlayer)
	isColFilled, colNum := board.isColFilled(board.currentPlayer)
	isDiagonalFilled, diagonalDirection := board.isDiagonalFilled(board.currentPlayer)

	isWinning := isRowFilled || isColFilled || isDiagonalFilled

	if isWinning {
		switch {
		case isRowFilled:
			board.isRowWinning = true
			board.winningOnRow = rowNum
		case isColFilled:
			board.isColWinning = true
			board.winningOnCol = colNum
		case isDiagonalFilled:
			board.isDiagonalWinning = true
			board.winningOnDiagonalDirection = diagonalDirection
		}
	}

	return isWinning
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

func (board *Board) isRowFilled(player Player) (bool, int) {
	for row := range board.size {
		filled := true
		for col := range board.size {
			if board.blocks[row][col].Player() != player {
				filled = false
				break
			}
		}

		if filled {
			return true, row
		}
	}

	return false, 0
}

func (board *Board) isColFilled(player Player) (bool, int) {
	for colNum := range board.size {
		filled := true
		for rowNum := range board.size {
			if board.blocks[rowNum][colNum].Player() != player {
				filled = false
				break
			}
		}

		if filled {
			return true, colNum
		}
	}

	return false, 0
}

func (board *Board) isDiagonalFilled(player Player) (bool, diagonalDirection) {
	filled := true
	// main diagonal
	for i := range board.size {
		if board.blocks[i][i].Player() != player {
			filled = false
			break
		}
	}

	if filled {
		return true, mainDiagonal
	}

	filled = true
	// anti diagonal
	for i := range board.size {
		if board.blocks[i][board.size-1-i].Player() != player {
			filled = false
			break
		}
	}
	if filled {
		return true, antiDiagonal
	}

	return false, unknownDiagonal
}

func (board *Board) setWinningLine() {
	if !board.IsGameOver() {
		return
	}

	if board.winningLine != nil {
		return
	}

	var line animation.Animation

	switch {
	case board.isRowWinning:
		onRow := board.winningOnRow
		startingPoint := animation.Point{
			X: float32(board.blocks[onRow][0].x) + linePadding,
			Y: float32(board.blocks[onRow][0].y) + blockSize/2,
		}
		endingPoint := animation.Point{
			X: float32(board.blocks[onRow][board.size-1].x) + blockSize - linePadding,
			Y: float32(board.blocks[onRow][board.size-1].y) + blockSize/2,
		}

		line = animation.NewLine(startingPoint, endingPoint, 5, 5, colorRed, true)

	case board.isColWinning:
		onCol := board.winningOnCol
		startingPoint := animation.Point{
			X: float32(board.blocks[0][onCol].x) + blockSize/2,
			Y: float32(board.blocks[0][onCol].y) + linePadding,
		}

		endingPoint := animation.Point{
			X: float32(board.blocks[board.size-1][onCol].x) + blockSize/2,
			Y: float32(board.blocks[board.size-1][onCol].y) + blockSize - linePadding,
		}

		line = animation.NewLine(startingPoint, endingPoint, 5, 5, colorRed, true)

	case board.isDiagonalWinning:
		switch board.winningOnDiagonalDirection {
		case mainDiagonal:
			startingPoint := animation.Point{
				X: float32(board.blocks[0][0].x) + linePadding,
				Y: float32(board.blocks[0][0].y) + linePadding,
			}

			endingPoint := animation.Point{
				X: float32(board.blocks[board.size-1][board.size-1].x) + blockSize - linePadding,
				Y: float32(board.blocks[board.size-1][board.size-1].y) + blockSize - linePadding,
			}

			line = animation.NewLine(startingPoint, endingPoint, 5, 5, colorRed, true)

		case antiDiagonal:
			startingPoint := animation.Point{
				X: float32(board.blocks[0][board.size-1].x) + blockSize - linePadding,
				Y: float32(board.blocks[0][board.size-1].y) + linePadding,
			}

			endingPoint := animation.Point{
				X: float32(board.blocks[board.size-1][0].x) + linePadding,
				Y: float32(board.blocks[board.size-1][0].y) + blockSize - linePadding,
			}

			line = animation.NewLine(startingPoint, endingPoint, 5, 5, colorRed, true)
		}
	}

	board.winningLine = line
}
