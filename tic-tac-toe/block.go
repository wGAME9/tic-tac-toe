package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	blockSize   = 100
	blockMargin = 4
)

type Block struct {
	x, y   int
	player Player
}

func NewBlock(rowNum, colNum int) *Block {
	x := colNum*blockSize + (colNum+1)*blockMargin
	y := rowNum*blockSize + (rowNum+1)*blockMargin

	return &Block{
		x:      x,
		y:      y,
		player: playerUnknown,
	}
}

func (block *Block) Draw(boardImage *ebiten.Image) {
	blockImage := ebiten.NewImage(blockSize, blockSize)
	blockImage.Fill(colorWhite)

	drawBlockOpt := &ebiten.DrawImageOptions{}
	drawBlockOpt.GeoM.Translate(float64(block.x), float64(block.y))
	boardImage.DrawImage(blockImage, drawBlockOpt)

	if block.IsUnknownPlayer() {
		return
	}

	// TODO: refactor this to make it more readable and maintainable
	// maybe making it a method of Player
	if block.Player() == playerX {
		// Draw a cross using vector line
		x, y := block.x, block.y
		padding := blockSize / 4
		vector.StrokeLine(
			boardImage,
			float32(x+padding), float32(y+padding),
			float32(x+blockSize-padding), float32(y+blockSize-padding),
			5,
			colorBlue,
			true,
		)

		vector.StrokeLine(
			boardImage,
			float32(x+blockSize-padding), float32(y+padding),
			float32(x+padding), float32(y+blockSize-padding),
			5,
			colorBlue,
			true,
		)
	} else {
		centerX := float32(block.x) + float32(blockSize/2)
		centerY := float32(block.y) + float32(blockSize/2)
		r := float32(blockSize) / 4
		vector.StrokeCircle(
			boardImage,
			centerX, centerY, r,
			5,
			block.player.Color(),
			true,
		)
	}

}

func (block *Block) BeingClicked(mouseX, mouseY int) bool {
	isInXRange := block.x < mouseX && mouseX < block.x+blockSize
	isInYRange := block.y < mouseY && mouseY < block.y+blockSize

	return isInXRange && isInYRange
}

func (block *Block) IsUnknownPlayer() bool {
	return block.player == playerUnknown
}

func (block *Block) CanSetPlayer() bool {
	return block.IsUnknownPlayer()
}

func (block *Block) SetPlayer(player Player) {
	block.player = player
}

func (block *Block) Player() Player {
	return block.player
}
