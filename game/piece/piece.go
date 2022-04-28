package piece

import "go_chess/game/constants"

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

func (p Piece) GetSuit() string {
	if p.Name == constants.BLACK_PAWN ||
		p.Name == constants.BLACK_ROOK ||
		p.Name == constants.BLACK_KNIGHT ||
		p.Name == constants.BLACK_BISHOP ||
		p.Name == constants.BLACK_KING ||
		p.Name == constants.BLACK_QUEEN {
		return "BLACK"
	}
	return "WHITE"
}

//func Build(name string, currentX int, currentY int) Piece{
//
//}
