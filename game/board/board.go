package board

import (
	"fmt"
	"go_chess/game/constants"
	"go_chess/game/piece"
)

type Game struct {
	Board [8][8]piece.Piece
}

func (g Game) InitializeBoard() Game {
	g.Board = [8][8]piece.Piece{
		{piece.Piece{Name: constants.BLACK_ROOK, CurrentPosition: "a8"},
			piece.Piece{Name: constants.BLACK_KNIGHT, CurrentPosition: "b8"},
			piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "c8"},
			piece.Piece{Name: constants.BLACK_QUEEN, CurrentPosition: "d8"},
			piece.Piece{Name: constants.BLACK_KING, CurrentPosition: "e8"},
			piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "f8"},
			piece.Piece{Name: constants.BLACK_KNIGHT, CurrentPosition: "g8"},
			piece.Piece{Name: constants.BLACK_ROOK, CurrentPosition: "h8"}},
		{piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "a7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "b7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "c7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "d7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "e7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "f7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "g7"},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "h7"}},
		{},
		{},
		{},
		{},
		{piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "a2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "b2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "c2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "d2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "e2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "f2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "g2"},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "h2"}},
		{piece.Piece{Name: constants.WHITE_ROOK, CurrentPosition: "a1"},
			piece.Piece{Name: constants.WHITE_KNIGHT, CurrentPosition: "b1"},
			piece.Piece{Name: constants.WHITE_BISHOP, CurrentPosition: "c1"},
			piece.Piece{Name: constants.WHITE_QUEEN, CurrentPosition: "d1"},
			piece.Piece{Name: constants.WHITE_KING, CurrentPosition: "e1"},
			piece.Piece{Name: constants.WHITE_BISHOP, CurrentPosition: "f1"},
			piece.Piece{Name: constants.WHITE_KNIGHT, CurrentPosition: "g1"},
			piece.Piece{Name: constants.WHITE_ROOK, CurrentPosition: "h1"}},
	}
	return g
}

func (g Game) InitializeEmptyBoard() Game {
	g.Board = [8][8]piece.Piece{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}
	return g
}

func (g Game) PrintBoard() {
	fmt.Println("***********   NEXT TURN   ***********")
	for i := 0; i < 8; i++ {
		fmt.Printf(" %d ", (8 - i))
		print("| ")
		print(g.Board[i][0].GetName())
		print(" | ")
		print(g.Board[i][1].GetName())
		print(" | ")
		print(g.Board[i][2].GetName())
		print(" | ")
		print(g.Board[i][3].GetName())
		print(" | ")
		print(g.Board[i][4].GetName())
		print(" | ")
		print(g.Board[i][5].GetName())
		print(" | ")
		print(g.Board[i][6].GetName())
		print(" | ")
		print(g.Board[i][7].GetName())
		print(" | ")
		println("")
	}
	println("     a   b   c   d   e   f   g   h")
	println("")
}
