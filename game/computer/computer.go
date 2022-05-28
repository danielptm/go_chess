package computer

import (
	"go_chess/game/board"
	"go_chess/game/constants"
)

//TODO: Loop through the board's black pieces
// Generate all possible moves for each piece
// Return a randomly selected piece and random move
// for that piece
func ComputerDecides(board board.Game) string {
	for i := 0; i < len(board.Board[0]); i++ {
		for j := 0; j < len(board.Board[1]); j++ {
			if board.Board[i][j].Name == constants.BLACK_KING {

			}
			if board.Board[i][j].Name == constants.BLACK_QUEEN {

			}
			if board.Board[i][j].Name == constants.BLACK_BISHOP {

			}
			if board.Board[i][j].Name == constants.BLACK_KNIGHT {

			}
			if board.Board[i][j].Name == constants.BLACK_ROOK {

			}
			if board.Board[i][j].Name == constants.BLACK_PAWN {

			}
		}
	}
	return ""
}
