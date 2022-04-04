package piece

import "strings"

type Piece struct {
	Name            string
	CurrentPosition string
	CurrentY        int
	CurrentX        int
	HasMoved        bool
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
