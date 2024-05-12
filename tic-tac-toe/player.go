package tictactoe

import "image/color"

type Player int

const (
	playerUnknown Player = iota
	playerX
	playerO
)

func (p Player) String() string {
	switch p {
	case playerX:
		return "X"
	case playerO:
		return "O"
	default:
		return "-"
	}
}

func (p Player) Color() color.Color {
	if p == playerX {
		return colorBlue
	}

	return colorRed
}
