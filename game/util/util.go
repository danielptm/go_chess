package util

import (
	"errors"
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
		if p.Name == constants.BLACK_PAWN {
			y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 2)

		}
		if p.Name == constants.WHITE_PAWN {
			y, x = DirectUpForSpaces(p.CurrentX, p.CurrentY, 2)

		}
		if IsInbounds(x, y) && b.Board[y][x].Name == "" {
			np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
			paths = append(paths, np.CurrentPosition)
		}
	}

	if p.Name == constants.BLACK_PAWN {
		y, x = DirectDownForSpaces(p.CurrentX, p.CurrentY, 1)
	}
	if p.Name == constants.WHITE_PAWN {
		y, x = DirectUpForSpaces(p.CurrentX, p.CurrentY, 1)
	}
	if IsInbounds(x, y) && b.Board[y][x].Name == "" {
		np := GetBoardPosition(piece.Piece{CurrentX: x, CurrentY: y})
		paths = append(paths, np.CurrentPosition)
	}

	leftTakeX, leftTakeY, rightTakeX, rightTakeY := -1, -1, -1, -1
	if p.Name == constants.BLACK_PAWN {
		leftTakeY, leftTakeX = DownLeftForSpaces(p.CurrentX, p.CurrentY, 1)
		rightTakeY, rightTakeX = DownRightForSpaces(p.CurrentX, p.CurrentY, 1)
	}

	if p.Name == constants.WHITE_PAWN {
		leftTakeY, leftTakeX = UpLeftForSpaces(p.CurrentX, p.CurrentY, 1)
		rightTakeY, rightTakeX = UpRightForSpaces(p.CurrentX, p.CurrentY, 1)
	}

	if !IsBlackPlayer(p) {
		if IsInbounds(leftTakeX, leftTakeY) && IsBlackPlayer(b.Board[leftTakeY][leftTakeX]) && !IsEmptySpace(b.Board[leftTakeY][leftTakeX]) {
			lefTakePos := GetBoardPosition(piece.Piece{CurrentX: leftTakeX, CurrentY: leftTakeY})
			paths = append(paths, lefTakePos.CurrentPosition)
		}

		if IsInbounds(rightTakeX, rightTakeY) && IsBlackPlayer(b.Board[rightTakeY][rightTakeX]) && !IsEmptySpace(b.Board[rightTakeY][rightTakeX]) {
			rightTakePos := GetBoardPosition(piece.Piece{CurrentX: rightTakeX, CurrentY: rightTakeY})
			paths = append(paths, rightTakePos.CurrentPosition)
		}
	}

	if IsBlackPlayer(p) {
		if IsInbounds(leftTakeX, leftTakeY) && !IsBlackPlayer(b.Board[leftTakeY][leftTakeX]) && !IsEmptySpace(b.Board[leftTakeY][leftTakeX]) {
			lefTakePos := GetBoardPosition(piece.Piece{CurrentX: leftTakeX, CurrentY: leftTakeY})
			paths = append(paths, lefTakePos.CurrentPosition)
		}

		if IsInbounds(rightTakeX, rightTakeY) && !IsBlackPlayer(b.Board[rightTakeY][rightTakeX]) && !IsEmptySpace(b.Board[rightTakeY][rightTakeX]) {
			rightTakePos := GetBoardPosition(piece.Piece{CurrentX: rightTakeX, CurrentY: rightTakeY})
			paths = append(paths, rightTakePos.CurrentPosition)
		}
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
	//paths = append(paths, UpBigLeftL(p, board)...)
	upBigLeftLPaths := UpBigLeftL(p, board)
	for _, v := range upBigLeftLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, UpBigRightL(p, board)...)
	upBigRightLPaths := UpBigRightL(p, board)
	for _, v := range upBigRightLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, UpSmallLeftL(p, board)...)
	upSmallLeftLPaths := UpSmallLeftL(p, board)
	for _, v := range upSmallLeftLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, UpSmallRightL(p, board)...)
	upSmallRightLPaths := UpSmallRightL(p, board)
	for _, v := range upSmallRightLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, RightBigUpL(p, board)...)
	rightBigUpLPaths := RightBigUpL(p, board)
	for _, v := range rightBigUpLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, RightBigDownL(p, board)...)
	rightBigDownLPaths := RightBigDownL(p, board)
	for _, v := range rightBigDownLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, RightSmallUpL(p, board)...)
	rightSmallUpLPaths := RightSmallUpL(p, board)
	for _, v := range rightSmallUpLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, RightSmallDownL(p, board)...)
	rightSmallDownLPaths := RightSmallDownL(p, board)
	for _, v := range rightSmallDownLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, DownBigLeftL(p, board)...)
	downBigLeftLPaths := DownBigLeftL(p, board)
	for _, v := range downBigLeftLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, DownBigRightL(p, board)...)
	downBigRightLPaths := DownBigRightL(p, board)
	for _, v := range downBigRightLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, DownSmallLeftL(p, board)...)
	downSmallLeftLPaths := DownSmallLeftL(p, board)
	for _, v := range downSmallLeftLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, DownSmallRightL(p, board)...)
	downSmallRightLPaths := DownSmallRightL(p, board)
	for _, v := range downSmallRightLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, LeftBigUpL(p, board)...)
	leftBigUpLPaths := LeftBigUpL(p, board)
	for _, v := range leftBigUpLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, LeftBigDownL(p, board)...)
	leftBigDownLPaths := LeftBigDownL(p, board)
	for _, v := range leftBigDownLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, LeftSmallUpL(p, board)...)
	leftSmallUpLPaths := LeftSmallUpL(p, board)
	for _, v := range leftSmallUpLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
	//paths = append(paths, LeftSmallDownL(p, board)...)
	leftSmallDownLPaths := LeftSmallDownL(p, board)
	for _, v := range leftSmallDownLPaths {
		if !Contains(v, paths) {
			paths = append(paths, v)
		}
	}
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
	p.HasMoved = true
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

func PlayMove(move string, b board.Game) board.Game {
	moveBits := strings.Split(move, ":")
	p := GetPieceFromPosition(moveBits[1], b)
	b, _ = PlacePiece(piece.Piece{}, moveBits[1], b)
	if GetPieceFromPosition(moveBits[2], b).Name != "" {
		b = TakePiece(moveBits[2], p, b)
	} else {
		b, _ = PlacePiece(p, moveBits[2], b)
	}
	return b
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

func TakePiece(pos string, p piece.Piece, b board.Game) board.Game {
	taken := GetPieceFromPosition(pos, b)
	b, _ = PlacePiece(piece.Piece{}, p.CurrentPosition, b)
	b.Cache = append(b.Cache, taken)
	b, _ = PlacePiece(p, pos, b)
	return b
}

func GenerateMoves(isBlack bool, board board.Game) ([][]string, error) {

	kingOptions := make([]string, 0)
	queenOptions := make([]string, 0)
	bishopOptions := make([]string, 0)
	knightOptions := make([]string, 0)
	rookOptions := make([]string, 0)
	pawnOptions := make([]string, 0)

	allMoves := [][]string{kingOptions, queenOptions, bishopOptions, knightOptions, rookOptions, pawnOptions}

	for i := 0; i < len(board.Board[0]); i++ {
		for j := 0; j < len(board.Board[1]); j++ {
			piece := ""
			if isBlack {
				piece = constants.BLACK_KING
			} else {
				piece = constants.WHITE_KING
			}
			if board.Board[i][j].Name == piece {
				kingOptions = GetKingPaths(board.Board[i][j], board)
				kingMoves := make([]string, 0)
				for _, v := range kingOptions {
					s := "king:" + board.Board[i][j].CurrentPosition + ":" + v
					kingMoves = append(kingMoves, s)
				}
				if len(kingMoves) > 0 {
					allMoves[0] = append(allMoves[0], kingMoves...)
				}
			}
			if isBlack {
				piece = constants.BLACK_QUEEN
			} else {
				piece = constants.WHITE_QUEEN
			}
			if board.Board[i][j].Name == piece {
				queenOptions = GetQueenPaths(board.Board[i][j], board)
				queenMoves := make([]string, 0)
				for _, v := range queenOptions {
					s := "queen:" + board.Board[i][j].CurrentPosition + ":" + v
					queenMoves = append(queenMoves, s)
				}
				if len(queenMoves) > 0 {
					allMoves[1] = append(allMoves[1], queenMoves...)
				}
			}
			if isBlack {
				piece = constants.BLACK_BISHOP
			} else {
				piece = constants.WHITE_BISHOP
			}
			if board.Board[i][j].Name == piece {
				bishopOptions = GetBishopPaths(board.Board[i][j], board)
				bishopMoves := make([]string, 0)
				for _, v := range bishopOptions {
					s := "bishop:" + board.Board[i][j].CurrentPosition + ":" + v
					bishopMoves = append(bishopMoves, s)
				}
				if len(bishopMoves) > 0 {
					allMoves[2] = append(allMoves[2], bishopMoves...)
				}
			}
			if isBlack {
				piece = constants.BLACK_KNIGHT
			} else {
				piece = constants.WHITE_KNIGHT
			}
			if board.Board[i][j].Name == piece {
				knightOptions = GetKnightPaths(board.Board[i][j], board)
				knightMoves := make([]string, 0)
				for _, v := range knightOptions {
					s := "knight:" + board.Board[i][j].CurrentPosition + ":" + v
					knightMoves = append(knightMoves, s)
				}
				if len(knightMoves) > 0 {
					allMoves[3] = append(allMoves[3], knightMoves...)
					print("1")
				}
			}
			if isBlack {
				piece = constants.BLACK_ROOK
			} else {
				piece = constants.WHITE_ROOK
			}
			if board.Board[i][j].Name == piece {
				rookOptions = GetRookPaths(board.Board[i][j], board)
				rookMoves := make([]string, 0)
				for _, v := range rookOptions {
					s := "rook:" + board.Board[i][j].CurrentPosition + ":" + v
					rookMoves = append(rookMoves, s)
				}
				if len(rookMoves) > 0 {
					allMoves[4] = append(allMoves[4], rookMoves...)
				}
			}
			if isBlack {
				piece = constants.BLACK_PAWN
			} else {
				piece = constants.WHITE_PAWN
			}
			if board.Board[i][j].Name == piece {
				pawnOptions = GetPawnPaths(board.Board[i][j], board)
				pawnMoves := make([]string, 0)
				for _, v := range pawnOptions {
					s := "pawn:" + board.Board[i][j].CurrentPosition + ":" + v
					pawnMoves = append(pawnMoves, s)
				}
				if len(pawnMoves) > 0 {
					allMoves[5] = append(allMoves[5], pawnMoves...)
				}
			}
		}
	}
	return allMoves, nil
}

func CheckIfHumanMoveIsValid(move string, b board.Game) (bool, error) {
	allMoves, _ := GenerateMoves(false, b)
	for _, v := range allMoves {
		if Contains(move, v) {
			return true, nil
		}
	}
	return false, errors.New("Invalid move input by human player.")
}

func IsBlackPlayer(p piece.Piece) bool {
	if p.Name == constants.BLACK_KING || p.Name == constants.BLACK_QUEEN || p.Name == constants.BLACK_KNIGHT || p.Name == constants.BLACK_BISHOP || p.Name == constants.BLACK_ROOK || p.Name == constants.BLACK_PAWN {
		return true
	}
	return false
}

func IsEmptySpace(p piece.Piece) bool {
	if p.Name == "" {
		return true
	}
	return false
}
