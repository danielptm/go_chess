package game

import (
	"fmt"
	"go_chess/constants"
)

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

func (g Game) PrintBoard() {
	fmt.Println("*********** NEXT TURN ***********")
	for i := 0; i < 8; i++ {
		fmt.Printf(" %d ", (8 - i))
		print("| ")
		print(g.Board[i][0])
		print(" | ")
		print(g.Board[i][1])
		print(" | ")
		print(g.Board[i][2])
		print(" | ")
		print(g.Board[i][3])
		print(" | ")
		print(g.Board[i][4])
		print(" | ")
		print(g.Board[i][5])
		print(" | ")
		print(g.Board[i][6])
		print(" | ")
		print(g.Board[i][7])
		print(" | ")
		println("")
	}
	println("     a   b   c   d   e   f   g   h")
	println("")
}
