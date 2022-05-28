package computer

import (
	"errors"
	"go_chess/game/board"
	"go_chess/game/constants"
	"go_chess/game/util"
	"math/rand"
)

//TODO: Loop through the board's black pieces
// Generate all possible moves for each piece
// Return a randomly selected piece and random move
// for that piece
func ComputerDecidesRandomly(board board.Game) (string, error) {

	kingOptions := make([]string, 0)
	queenOptions := make([]string, 0)
	bishopOptions := make([]string, 0)
	knightOptions := make([]string, 0)
	rookOptions := make([]string, 0)
	pawnOptions := make([]string, 0)

	kingMoves := make([]string, 0)
	queenMoves := make([]string, 0)
	bishopMoves := make([]string, 0)
	knightMoves := make([]string, 0)
	rookMoves := make([]string, 0)
	pawnMoves := make([]string, 0)

	for i := 0; i < len(board.Board[0]); i++ {
		for j := 0; j < len(board.Board[1]); j++ {
			if board.Board[i][j].Name == constants.BLACK_KING {
				kingOptions = util.GetKingPaths(board.Board[i][j], board)
				for i, v := range kingOptions {
					s := "king:" + board.Board[i][j].CurrentPosition + ":" + v
					kingMoves = append(kingMoves, s)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_QUEEN {
				queenOptions = util.GetQueenPaths(board.Board[i][j], board)
				for i, v := range queenOptions {
					s := "queen:" + board.Board[i][j].CurrentPosition + ":" + v
					queenMoves = append(queenMoves, s)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_BISHOP {
				bishopOptions = util.GetBishopPaths(board.Board[i][j], board)
				for i, v := range bishopOptions {
					s := "bishop:" + board.Board[i][j].CurrentPosition + ":" + v
					bishopMoves = append(bishopMoves, s)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_KNIGHT {
				knightOptions = util.GetKingPaths(board.Board[i][j], board)
				for i, v := range knightOptions {
					s := "knight:" + board.Board[i][j].CurrentPosition + ":" + v
					knightMoves = append(knightMoves, s)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_ROOK {
				rookOptions = util.GetKingPaths(board.Board[i][j], board)
				for i, v := range rookOptions {
					s := "rook:" + board.Board[i][j].CurrentPosition + ":" + v
					rookMoves = append(rookMoves, s)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_PAWN {
				pawnOptions = util.GetKingPaths(board.Board[i][j], board)
				for i, v := range pawnOptions {
					s := "pawn:" + board.Board[i][j].CurrentPosition + ":" + v
					pawnMoves = append(pawnMoves, s)
				}
			}
		}
	}
	max := generateIntMax(kingOptions, queenOptions, bishopOptions, knightOptions, rookOptions, pawnOptions)
	min := 0
	pieceR := rand.Intn(max-min+1) + min

	switch pieceR {
	case 0:
		max := len(kingMoves)
		min := 0
		pieceR := rand.Intn(max-min+1) + min
		return kingMoves[pieceR], nil
	case 1:
		max := len(queenMoves)
		min := 0
		pieceR := rand.Intn(max-min+1) + min
		return queenMoves[pieceR], nil
	case 2:
		max := len(bishopMoves)
		min := 0
		pieceR := rand.Intn(max-min+1) + min
		return bishopMoves[pieceR], nil
	case 3:
		max := len(knightMoves)
		min := 0
		pieceR := rand.Intn(max-min+1) + min
		return knightMoves[pieceR], nil
	case 4:
		max := len(rookMoves)
		min := 0
		pieceR := rand.Intn(max-min+1) + min
		return rookMoves[pieceR], nil
	case 5:
		max := len(pawnMoves)
		min := 0
		pieceR := rand.Intn(max-min+1) + min
		return pawnMoves[pieceR], nil
	}
	return "-99", errors.New("There was a problem generating a random move for the computer")
}

func generateIntMax(
	kingOptions []string,
	queenOptions []string,
	bishopOptions []string,
	knightOptions []string,
	rookOptions []string,
	pawnOptions []string) int {

	max := 0
	if len(kingOptions) > 0 {
		max += 1
	}
	if len(queenOptions) > 0 {
		max += 1
	}
	if len(bishopOptions) > 0 {
		max += 1
	}
	if len(knightOptions) > 0 {
		max += 1
	}
	if len(rookOptions) > 0 {
		max += 1
	}
	if len(pawnOptions) > 0 {
		max += 1
	}
	return max
}
