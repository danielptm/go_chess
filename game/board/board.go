package board

import (
	"fmt"
	"go_chess/game/constants"
	"go_chess/game/piece"
	"strconv"
	"strings"
)

type Game struct {
	Board          [8][8]piece.Piece
	Cache          []piece.Piece
	HumanIsInCheck bool
}

func (g Game) InitializeBoard() Game {
	g.Cache = make([]piece.Piece, 0)
	g.Board = [8][8]piece.Piece{
		{piece.Piece{Name: constants.BLACK_ROOK, CurrentPosition: "a8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_KNIGHT, CurrentPosition: "b8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "c8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_QUEEN, CurrentPosition: "d8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_KING, CurrentPosition: "e8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "f8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_KNIGHT, CurrentPosition: "g8", HasMoved: false},
			piece.Piece{Name: constants.BLACK_ROOK, CurrentPosition: "h8", HasMoved: false}},
		{piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "a7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "b7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "c7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "d7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "e7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "f7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "g7", HasMoved: false},
			piece.Piece{Name: constants.BLACK_PAWN, CurrentPosition: "h7", HasMoved: false}},
		{},
		{},
		{},
		{},
		{piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "a2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "b2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "c2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "d2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "e2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "f2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "g2", HasMoved: false},
			piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "h2", HasMoved: false}},
		{piece.Piece{Name: constants.WHITE_ROOK, CurrentPosition: "a1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_KNIGHT, CurrentPosition: "b1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_BISHOP, CurrentPosition: "c1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_QUEEN, CurrentPosition: "d1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_KING, CurrentPosition: "e1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_BISHOP, CurrentPosition: "f1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_KNIGHT, CurrentPosition: "g1", HasMoved: false},
			piece.Piece{Name: constants.WHITE_ROOK, CurrentPosition: "h1", HasMoved: false}},
	}

	for i, _ := range g.Board {
		for j, _ := range g.Board[i] {
			if g.Board[i][j].Name != "" {
				p := GetCoordinates(g.Board[i][j])
				g.Board[i][j] = p

			}
		}
	}

	return g
}

//TODO: This is duplicate code from util... Due to circle dep, had to put it here.
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

func (g Game) InitializeEmptyBoard() Game {
	g.Cache = make([]piece.Piece, 0)
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

func (g Game) PrintCache() {
	for i := 0; i < len(g.Cache); i++ {
		println(GetTextForPiece(g.Cache[i]))
	}
}

func (g Game) PrintBoardDetail() {
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			if g.Board[i][j].Name != "" {
				println(g.Board[i][j].CurrentPosition + ": " + GetTextForPiece(g.Board[i][j]))
			}
		}
	}
}

func (g Game) PrintBoard() {
	fmt.Println("***********   NEXT TURN   ***********")
	fmt.Println("Human is in check: " + strconv.FormatBool(g.HumanIsInCheck))
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

func GetTextForPiece(p piece.Piece) string {
	if p.Name == constants.WHITE_KING {
		return "White king"
	}
	if p.Name == constants.WHITE_QUEEN {
		return "White queen"
	}
	if p.Name == constants.WHITE_BISHOP {
		return "White bishop"
	}
	if p.Name == constants.WHITE_KNIGHT {
		return "White knight"
	}
	if p.Name == constants.WHITE_ROOK {
		return "White rook"
	}
	if p.Name == constants.WHITE_PAWN {
		return "White pawn"
	}
	if p.Name == constants.BLACK_KING {
		return "Black king"
	}
	if p.Name == constants.BLACK_QUEEN {
		return "Black queen"
	}
	if p.Name == constants.BLACK_BISHOP {
		return "Black bishop"
	}
	if p.Name == constants.BLACK_KNIGHT {
		return "Black knight"
	}
	if p.Name == constants.BLACK_ROOK {
		return "Black rook"
	}
	if p.Name == constants.BLACK_PAWN {
		return "Black pawn"
	}
	return "N/A"
}
