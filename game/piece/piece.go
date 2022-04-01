package piece

import "strings"

type Piece struct {
	Name            string
	CurrentPosition string
	currentY        int
	currentX        int
	PossibleMoves   []string
}

func (p Piece) GetName() string {
	if p.Name == "" {
		return " "
	} else {
		return p.Name
	}
}

//TODO: Add a unique and possible move to the possibleMoves array
func (p *Piece) AddMove() {

}

//TODO: Generate a theoretically possible move. The board will determine if it is ok.
func (p Piece) GenerateMoves() []string {
	if strings.Contains(p.Name, "KING") {

	}
	if strings.Contains(p.Name, "QUEEN") {

	}
	if strings.Contains(p.Name, "KNIGHT") {

	}
	if strings.Contains(p.Name, "ROOK") {

	}
	if strings.Contains(p.Name, "BISHOP") {

	}
	if strings.Contains(p.Name, "PAWN") {

	}
	return make([]string, 5)
}

//func KingMoves(y int, x int) [][]int {
//	moves := make([][]int, 2)
//	append(moves, [y + 1, x])
//}

func (p *Piece) GetCoordinates() (int, int) {
	split := strings.Split(p.CurrentPosition, "")

	switch split[0] {
	case "a":
		p.currentX = 0
	case "b":
		p.currentX = 1
	case "c":
		p.currentX = 2
	case "d":
		p.currentX = 3
	case "e":
		p.currentX = 4
	case "f":
		p.currentX = 5
	case "g":
		p.currentX = 6
	case "h":
		p.currentX = 7
	}
	switch split[1] {
	case "8":
		p.currentY = 0
	case "7":
		p.currentY = 1
	case "6":
		p.currentY = 2
	case "5":
		p.currentY = 3
	case "4":
		p.currentY = 4
	case "3":
		p.currentY = 5
	case "2":
		p.currentY = 6
	case "1":
		p.currentY = 7
	}
	return p.currentX, p.currentY
}

func (p *Piece) GetBoardPosition() string {
	postion := ""
	switch p.currentX {
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

	switch p.currentY {
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
	return p.CurrentPosition
}
