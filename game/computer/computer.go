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
func ComputerDecides(board board.Game) (string, error) {

	kingMoves := make([]string, 0)
	queenMoves := make([]string, 0)
	bishopMoves := make([]string, 0)
	knightMoves := make([]string, 0)
	rookMoves := make([]string, 0)
	pawnMoves := make([]string, 0)

	for i := 0; i < len(board.Board[0]); i++ {
		for j := 0; j < len(board.Board[1]); j++ {
			if board.Board[i][j].Name == constants.BLACK_KING {
				kingMoves = util.GetKingPaths(board.Board[i][j])
			}
			if board.Board[i][j].Name == constants.BLACK_QUEEN {
				queenMoves = util.GetQueenPaths(board.Board[i][j], board)

			}
			if board.Board[i][j].Name == constants.BLACK_BISHOP {
				bishopMoves = util.GetBishopPaths(board.Board[i][j], board)

			}
			if board.Board[i][j].Name == constants.BLACK_KNIGHT {
				knightMoves = util.GetKingPaths(board.Board[i][j])

			}
			if board.Board[i][j].Name == constants.BLACK_ROOK {
				rookMoves = util.GetKingPaths(board.Board[i][j])
			}
			if board.Board[i][j].Name == constants.BLACK_PAWN {
				pawnMoves = util.GetKingPaths(board.Board[i][j])
			}
		}
	}
	max := 5
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
