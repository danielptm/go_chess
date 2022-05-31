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

//TODO: Add test for when there is a piece in the pawns way
func GetPawnPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := -1, -1
	if !p.HasMoved {
		if strings.Contains(p.Name, "black") {
			y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
		} else {
			y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
		}
		if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}

	if strings.Contains(p.Name, "black") {
		y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	} else {
		y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	}
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	return paths
}

func GetRookPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	paths = append(paths, GetDirectUpPaths(p, b)...)
	paths = append(paths, GetRightPaths(p, b)...)
	paths = append(paths, GetDownPaths(p, b)...)
	paths = append(paths, GetLeftPaths(p, b)...)
	return paths
}

//Does not take into account castling
//TODO: Add test for when something is in kings way
//TODO: Add test for when KING moves into check, should not be able to do this.
func GetKingPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)

	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)

	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = UpRightForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DirectRightForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DownRightForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DownLeftForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = DirectLeftForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	y, x = UpLeftForSpaces(p.CurrentX, p.CurrentY, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	return paths
}

func GetQueenPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	paths = append(paths, GetDirectUpPaths(p, b)...)
	paths = append(paths, GetUpRightPaths(p, b)...)
	paths = append(paths, GetRightPaths(p, b)...)
	paths = append(paths, GetDownRightPaths(p, b)...)
	paths = append(paths, GetDownPaths(p, b)...)
	paths = append(paths, GetDownLeftPaths(p, b)...)
	paths = append(paths, GetLeftPaths(p, b)...)
	paths = append(paths, GetUpLeftPaths(p, b)...)
	return paths
}

func GetDirectUpPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	upMax := p.CurrentY

	for i := 1; i <= upMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectUpForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetUpRightPaths(p piece.Piece, b board.Game) []string {
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
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetRightPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	rightMax := 7 - p.CurrentX

	for i := 1; i <= rightMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectRightForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetDownRightPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	downRightMax := -1

	//TODO: Need to add test for downRightMax is equal to p.currentY
	if (7 - p.CurrentX) < (7 - p.CurrentY) {
		downRightMax = p.CurrentX
	} else {
		downRightMax = p.CurrentY
	}

	for i := 1; i <= downRightMax; i++ {
		np := GetCoordinates(p)
		y, x := DownRightForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetDownPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)

	downMax := 7 - p.CurrentY

	for i := 1; i <= downMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectDownForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetDownLeftPaths(p piece.Piece, b board.Game) []string {
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
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetLeftPaths(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)

	leftMax := p.CurrentX

	for i := 1; i <= leftMax; i++ {
		np := GetCoordinates(p)
		y, x := DirectLeftForSpaces(np.CurrentX, np.CurrentY, i)
		if IsInbounds(x, y) {
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

//TODO: Take the board as another param. Search the board
// and find all paths. The furthest path is either the edge of the board
// or a piece. If the encountered piece is a different suit, then it
// returns that coordinate as a path. If it is the same suit, then it does not include it
// Do this for all of the GetPaths functions
func GetUpLeftPaths(p piece.Piece, b board.Game) []string {
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
			if IsSameColor(p.Name, b.Board[y][x].Name) {
				break
			}
			np = piece.Piece{Name: np.Name, CurrentX: x, CurrentY: y}
			np = GetBoardPosition(np)
			pos := np.CurrentPosition
			paths = append(paths, pos)
			if b.Board[y][x].Name != "" {
				break
			}
		}
	}
	return paths
}

func GetBishopPaths(p piece.Piece, g board.Game) []string {
	paths := make([]string, 0)
	paths = append(paths, GetUpLeftPaths(p, g)...)
	paths = append(paths, GetUpRightPaths(p, g)...)
	paths = append(paths, GetDownRightPaths(p, g)...)
	paths = append(paths, GetDownLeftPaths(p, g)...)
	return paths
}

func GetKnightPaths(p piece.Piece, board board.Game) []string {
	paths := make([]string, 0)
	paths = append(paths, UpBigLeftL(p, board)...)
	paths = append(paths, UpBigRightL(p, board)...)
	paths = append(paths, UpSmallLeftL(p, board)...)
	paths = append(paths, UpSmallRightL(p, board)...)
	paths = append(paths, RightBigUpL(p, board)...)
	paths = append(paths, RightBigDownL(p, board)...)
	paths = append(paths, RightSmallUpL(p, board)...)
	paths = append(paths, RightSmallDownL(p, board)...)
	paths = append(paths, DownBigLeftL(p, board)...)
	paths = append(paths, DownBigRightL(p, board)...)
	paths = append(paths, DownSmallLeftL(p, board)...)
	paths = append(paths, DownSmallRightL(p, board)...)
	paths = append(paths, LeftBigUpL(p, board)...)
	paths = append(paths, LeftBigDownL(p, board)...)
	paths = append(paths, LeftSmallUpL(p, board)...)
	paths = append(paths, LeftSmallDownL(p, board)...)
	return paths
}

//***
func UpBigLeftL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func UpBigRightL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func UpSmallLeftL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func UpSmallRightL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

//***
func RightBigUpL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func RightBigDownL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func RightSmallUpL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectRightForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func RightSmallDownL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectRightForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

//***
func DownBigLeftL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func DownBigRightL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectRightForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func DownSmallLeftL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func DownSmallRightL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectRightForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

//***
func LeftBigUpL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func LeftBigDownL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	y, x = DirectLeftForSpaces(x, y, 2)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func LeftSmallUpL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func LeftSmallDownL(p piece.Piece, b board.Game) []string {
	paths := make([]string, 0)
	y, x := DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)
	y, x = DirectLeftForSpaces(x, y, 1)
	if IsInbounds(x, y) && !IsSameColor(p.Name, b.Board[y][x].Name) {
		np := piece.Piece{CurrentX: x, CurrentY: y, Name: constants.BLACK_KNIGHT}
		paths = append(paths, GetBoardPosition(np).CurrentPosition)
	}
	return paths
}

func PlacePiece(p piece.Piece, pos string, g board.Game) (board.Game, piece.Piece) {
	p.CurrentPosition = pos
	np := GetCoordinates(p)
	np = GetBoardPosition(np)
	g.Board[np.CurrentY][np.CurrentX] = np
	return g, np
}

func IsInbounds(x int, y int) bool {
	if x >= 0 && x <= 7 && y >= 0 && y <= 7 {
		return true
	}
	return false
}

//TODO Get a a piece at the position by giving pram such as "a8"
func GetPieceFromPosition(pos string, g board.Game) piece.Piece {
	split := strings.Split(pos, "")
	CurrentX := -1
	CurrentY := -1
	switch split[0] {
	case "a":
		CurrentX = 0
	case "b":
		CurrentX = 1
	case "c":
		CurrentX = 2
	case "d":
		CurrentX = 3
	case "e":
		CurrentX = 4
	case "f":
		CurrentX = 5
	case "g":
		CurrentX = 6
	case "h":
		CurrentX = 7
	}
	switch split[1] {
	case "8":
		CurrentY = 0
	case "7":
		CurrentY = 1
	case "6":
		CurrentY = 2
	case "5":
		CurrentY = 3
	case "4":
		CurrentY = 4
	case "3":
		CurrentY = 5
	case "2":
		CurrentY = 6
	case "1":
		CurrentY = 7
	}
	return g.Board[CurrentY][CurrentX]
}

func IsSameColor(p1Name string, p2Name string) bool {
	whites := []string{constants.WHITE_PAWN, constants.WHITE_ROOK, constants.WHITE_KNIGHT, constants.WHITE_BISHOP, constants.WHITE_QUEEN, constants.WHITE_KING}
	blacks := []string{constants.BLACK_PAWN, constants.BLACK_ROOK, constants.BLACK_KNIGHT, constants.BLACK_BISHOP, constants.BLACK_QUEEN, constants.BLACK_KING}
	if Contains(p1Name, whites) && Contains(p2Name, whites) {
		return true
	} else if Contains(p1Name, blacks) && Contains(p2Name, blacks) {
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
