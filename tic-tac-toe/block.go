package tictactoe

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontSize    = 48.0
	blockSize   = 100
	blockMargin = 4
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

type Block struct {
	x, y   int
	player Player
}

func NewBlock(i, j int) *Block {
	x := i*blockSize + (i+1)*blockMargin
	y := j*blockSize + (j+1)*blockMargin

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

	drawPlayerOpt := &text.DrawOptions{}
	playerX := float64(block.x) + float64(blockSize/2)
	playerY := float64(block.y) + float64(blockSize/2)
	drawPlayerOpt.GeoM.Translate(playerX, playerY)
	drawPlayerOpt.ColorScale.ScaleWithColor(block.player.Color())
	drawPlayerOpt.PrimaryAlign = text.AlignCenter
	drawPlayerOpt.SecondaryAlign = text.AlignCenter
	text.Draw(
		boardImage,
		block.player.String(),
		&text.GoTextFace{
			Source: mplusFaceSource,
			Size:   fontSize,
		},
		drawPlayerOpt,
	)
}

func (block *Block) BeingClicked(mouseX, mouseY int) bool {
	isInXRange := block.x <= mouseX && mouseX < block.x+blockSize
	isInYRange := block.y <= mouseY && mouseY < block.y+blockSize

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
