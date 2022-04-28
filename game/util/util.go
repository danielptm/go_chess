package util

import (
	"go_chess/game/board"
	"go_chess/game/constants"
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
		if IsInbounds(x, y) {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}

	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	return paths
}

func GetRookPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	for i := 1; i <= p.CurrentY; i++ {
		y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, i)
		if IsInbounds(x, y) {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}

	maxRight := 7 - p.CurrentX
	for i := 1; i <= maxRight; i++ {
		y, x := DirectRightForSpaces(p.CurrentX, p.CurrentY, i)
		if IsInbounds(x, y) {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}

	maxDown := 7 - p.CurrentY

	for i := 1; i <= maxDown; i++ {
		y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, i)
		if IsInbounds(x, y) {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}

	maxLeft := p.CurrentX

	for i := 1; i <= maxLeft; i++ {
		y, x := DirectLeftForSpaces(p.CurrentX, p.CurrentY, i)
		if IsInbounds(x, y) {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}
	return paths
}

//Does not take into account castling
func GetKingPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)

	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = UpRightForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DirectRightForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DownRightForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DownLeftForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DirectLeftForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = UpLeftForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	return paths
}

func GetQueenPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	paths = append(paths, GetDirectUpPaths(p)...)
	paths = append(paths, GetUpRightPaths(p)...)
	paths = append(paths, GetRightPaths(p)...)
	paths = append(paths, GetDownRightPaths(p)...)
	paths = append(paths, GetDownPaths(p)...)
	paths = append(paths, GetDownLeftPaths(p)...)
	paths = append(paths, GetLeftPaths(p)...)
	paths = append(paths, GetUpLeftPaths(p)...)
	return paths
}

//TODO: Take the board as another param. Search up the board
// and find all paths. The furthest path is either the edge of the board
// or a piece. If the encountered piece is a different suit, then it
// returns that coordinate as a path. If it is the same suit, then it does
// it does not include it
// Do this for all of the GetPaths functions
func GetDirectUpPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	upMax := p.CurrentY

	for i := 1; i <= upMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectUpForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetUpRightPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	upRightMax := -1
	if p.CurrentY > (7 - p.CurrentX) {
		upRightMax = p.CurrentX
	} else {
		upRightMax = p.CurrentY
	}

	for i := 1; i <= upRightMax; i++ {
		np := GetCoordinates(p)
		y, x := UpRightForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetRightPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	rightMax := 7 - p.CurrentX

	for i := 1; i <= rightMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectRightForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetDownRightPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	downRightMax := -1

	if (7 - p.CurrentX) > (7 - p.CurrentY) {
		downRightMax = p.CurrentX
	} else {
		downRightMax = p.CurrentY
	}

	for i := 1; i <= downRightMax; i++ {
		np := GetCoordinates(p)
		y, x := DownRightForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetDownPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	downMax := 7 - p.CurrentY

	for i := 1; i <= downMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectDownForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetDownLeftPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	downLeftMax := -1
	if p.CurrentX > (7 - p.CurrentY) {
		downLeftMax = p.CurrentY
	} else {
		downLeftMax = p.CurrentX
	}

	for i := 1; i <= downLeftMax; i++ {
		np := GetCoordinates(p)
		y, x := DownLeftForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetLeftPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	leftMax := p.CurrentX

	for i := 1; i <= leftMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectLeftForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetUpLeftPaths(p piece.Piece) []string {
	paths := make([]string, 0)

	upLeftMax := -1
	if p.CurrentY > p.CurrentX {
		upLeftMax = p.CurrentX
	} else {
		upLeftMax = p.CurrentY
	}

	for i := 1; i <= upLeftMax; i++ {
		np := GetCoordinates(p)
		y, x := UpLeftForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
		}
	}
	return paths
}

func GetBishopPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	paths = append(paths, GetUpLeftPaths(p)...)
	paths = append(paths, GetUpRightPaths(p)...)
	paths = append(paths, GetDownRightPaths(p)...)
	paths = append(paths, GetDownLeftPaths(p)...)
	return paths
}

func GetKnightPaths(p piece.Piece) []string {
	paths := make([]string, 0)
	paths = append(paths, UpBigLeftL(p)...)
	paths = append(paths, UpBigRightL(p)...)
	paths = append(paths, UpSmallLeftL(p)...)
	paths = append(paths, UpSmallRightL(p)...)
	paths = append(paths, RightBigUpL(p)...)
	paths = append(paths, RightBigDownL(p)...)
	paths = append(paths, RightSmallUpL(p)...)
	paths = append(paths, RightSmallDownL(p)...)
	paths = append(paths, DownBigLeftL(p)...)
	paths = append(paths, DownBigRightL(p)...)
	paths = append(paths, DownSmallLeftL(p)...)
	paths = append(paths, DownSmallRightL(p)...)
	paths = append(paths, LeftBigUpL(p)...)
	paths = append(paths, LeftBigDownL(p)...)
	paths = append(paths, LeftSmallUpL(p)...)
	paths = append(paths, LeftSmallDownL(p)...)
	return paths
}

//***
func UpBigLeftL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func UpBigRightL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func UpSmallLeftL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func UpSmallRightL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

//***
func RightBigUpL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func RightBigDownL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func RightSmallUpL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectRightForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func RightSmallDownL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectRightForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

//***
func DownBigLeftL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func DownBigRightL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func DownSmallLeftL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func DownSmallRightL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectRightForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

//***
func LeftBigUpL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func LeftBigDownL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func LeftSmallUpL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func LeftSmallDownL(p piece.Piece) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func PlacePiece(p piece.Piece, pos string, g board.Game) board.Game {
	p.CurrentPosition = pos
	np := GetCoordinates(p)
	np = GetBoardPosition(np)
	g.Board[np.CurrentY][np.CurrentX] = np
	return g
}

func IsInbounds(x int, y int) bool {
	if x >= 0 && x <= 7 && y >= 0 && y <= 7 {
		return true
	}
	return false
}

func Contains(item string, a []string) bool {
	for _, s := range a {
		if s == item {
			return true
		}
	}
	return false
}
