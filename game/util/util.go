package util

import (
	"go_chess/game/piece"
	"strings"
)

func DirectUpForSpaces(x int, y int, spaces int) (int, int) {
	return y - spaces, x
}

func DirectDownForSpaces(x int, y int, spaces int) (int, int) {
	return y + spaces, x
}

func DirectRightForSpaces(x int, y int, spaces int) (int, int) {
	return y, x + spaces
}

func DirectLeftForSpaces(x int, y int, spaces int) (int, int) {
	return y, x - spaces
}

func UpRightForSpaces(x int, y int, spaces int) (int, int) {
	return y - spaces, x + spaces
}

func DownRightForSpaces(x int, y int, spaces int) (int, int) {
	return y + spaces, x + spaces
}

func UpLeftForSpaces(x int, y int, spaces int) (int, int) {
	return y - spaces, x - spaces
}

func DownLeftForSpaces(x int, y int, spaces int) (int, int) {
	return y + spaces, x - spaces
}

/*
	Returns x,y as int
*/
func GetCoordinates(p piece.Piece) piece.Piece {
	split := strings.Split(p.CurrentPosition, "")

	switch split[0] {
	case "a":
		p.CurrentX = 0
	case "b":
		p.CurrentX = 1
	case "c":
		p.CurrentX = 2
	case "d":
		p.CurrentX = 3
	case "e":
		p.CurrentX = 4
	case "f":
		p.CurrentX = 5
	case "g":
		p.CurrentX = 6
	case "h":
		p.CurrentX = 7
	}
	switch split[1] {
	case "8":
		p.CurrentY = 0
	case "7":
		p.CurrentY = 1
	case "6":
		p.CurrentY = 2
	case "5":
		p.CurrentY = 3
	case "4":
		p.CurrentY = 4
	case "3":
		p.CurrentY = 5
	case "2":
		p.CurrentY = 6
	case "1":
		p.CurrentY = 7
	}
	return p
}

/*
Returns x,y as int
*/
func GetBoardPosition(p piece.Piece) piece.Piece {
	postion := ""
	switch p.CurrentX {
	case 0:
		postion += "a"
	case 1:
		postion += "b"
	case 2:
		postion += "c"
	case 3:
		postion += "d"
	case 4:
		postion += "e"
	case 5:
		postion += "f"
	case 6:
		postion += "g"
	case 7:
		postion += "h"
	}

	switch p.CurrentY {
	case 0:
		postion += "8"
	case 1:
		postion += "7"
	case 2:
		postion += "6"
	case 3:
		postion += "5"
	case 4:
		postion += "4"
	case 5:
		postion += "3"
	case 6:
		postion += "2"
	case 7:
		postion += "1"
	}
	p.CurrentPosition = postion
	return p
}

func GetPawnPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	if !p.HasMoved {
		y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
	paths = append(paths, np.CurrentPosition)

	return paths
}

func GetRookPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	for i := 1; i <= p.CurrentY; i++ {
		y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, i)
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	maxRight := 7 - p.CurrentX
	for i := 1; i <= maxRight; i++ {
		y, x := DirectRightForSpaces(p.CurrentX, p.CurrentY, i)
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	maxDown := 7 - p.CurrentY

	for i := 1; i <= maxDown; i++ {
		y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, i)
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	maxLeft := p.CurrentX

	for i := 1; i <= maxLeft; i++ {
		y, x := DirectLeftForSpaces(p.CurrentX, p.CurrentY, i)
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	return paths
}

func GetBishopPaths(p piece.Piece) []string {
	return nil
}

func GetKnightPaths(p piece.Piece) []string {
	return nil
}

//Does not take into account castling
func GetKingPaths(p piece.Piece) []string {
	return nil
}

func GetQueenPaths(p piece.Piece) []string {
	return nil
}

func Contains(item string, a []string) bool {
	for _, s := range a {
		if s == item {
			return true
		}
	}
	return false
}
