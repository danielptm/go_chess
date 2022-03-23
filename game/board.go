package game

import "go_chess/constants"

type Game struct {
	Board [8][8]string
}

func (g Game) InitializeBoard() Game {
	g.Board = [8][8]string{
		{constants.BLACK_ROOK, constants.BLACK_KNIGHT, constants.BLACK_BISHOP, constants.BLACK_QUEEN, constants.BLACK_KING, constants.BLACK_BISHOP, constants.BLACK_KNIGHT, constants.BLACK_ROOK},
		{constants.BLACK_PAWN, constants.BLACK_PAWN, constants.BLACK_PAWN, constants.BLACK_PAWN, constants.BLACK_PAWN, constants.BLACK_PAWN, constants.BLACK_PAWN, constants.BLACK_PAWN},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{constants.WHITE_PAWN, constants.WHITE_PAWN, constants.WHITE_PAWN, constants.WHITE_PAWN, constants.WHITE_PAWN, constants.WHITE_PAWN, constants.WHITE_PAWN, constants.WHITE_PAWN},
		{constants.WHITE_ROOK, constants.WHITE_KNIGHT, constants.WHITE_BISHOP, constants.WHITE_QUEEN, constants.WHITE_KING, constants.WHITE_BISHOP, constants.WHITE_KNIGHT, constants.WHITE_ROOK},
	}
	return g
}
