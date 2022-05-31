package computer

import (
	"errors"
	"go_chess/game/board"
	"go_chess/game/constants"
	"go_chess/game/util"
	"math/rand"
	"time"
)

func ComputerRandomlyDecides(board board.Game) string {
	res, _ := generateMoves(board)
	move, _ := chooseRandomMoveForRandomPiece(res)
	return move
}

//TODO: Add better unit test for this
func generateMoves(board board.Game) ([][]string, error) {

	kingOptions := make([]string, 0)
	queenOptions := make([]string, 0)
	bishopOptions := make([]string, 0)
	knightOptions := make([]string, 0)
	rookOptions := make([]string, 0)
	pawnOptions := make([]string, 0)

	allMoves := make([][]string, 0)

	for i := 0; i < len(board.Board[0]); i++ {
		for j := 0; j < len(board.Board[1]); j++ {
			if board.Board[i][j].Name == constants.BLACK_KING {
				kingOptions = util.GetKingPaths(board.Board[i][j], board)
				kingMoves := make([]string, 0)
				for _, v := range kingOptions {
					s := "king:" + board.Board[i][j].CurrentPosition + ":" + v
					kingMoves = append(kingMoves, s)
				}
				if len(kingMoves) > 0 {
					allMoves = append(allMoves, kingMoves)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_QUEEN {
				queenOptions = util.GetQueenPaths(board.Board[i][j], board)
				queenMoves := make([]string, 0)
				for _, v := range queenOptions {
					s := "queen:" + board.Board[i][j].CurrentPosition + ":" + v
					queenMoves = append(queenMoves, s)
				}
				if len(queenMoves) > 0 {
					allMoves = append(allMoves, queenMoves)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_BISHOP {
				bishopOptions = util.GetBishopPaths(board.Board[i][j], board)
				bishopMoves := make([]string, 0)
				for _, v := range bishopOptions {
					s := "bishop:" + board.Board[i][j].CurrentPosition + ":" + v
					bishopMoves = append(bishopMoves, s)
				}
				if len(bishopMoves) > 0 {
					allMoves = append(allMoves, bishopMoves)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_KNIGHT {
				knightOptions = util.GetKnightPaths(board.Board[i][j], board)
				knightMoves := make([]string, 0)
				for _, v := range knightOptions {
					s := "knight:" + board.Board[i][j].CurrentPosition + ":" + v
					knightMoves = append(knightMoves, s)
				}
				if len(knightMoves) > 0 {
					allMoves = append(allMoves, knightMoves)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_ROOK {
				rookOptions = util.GetRookPaths(board.Board[i][j], board)
				rookMoves := make([]string, 0)
				for _, v := range rookOptions {
					s := "rook:" + board.Board[i][j].CurrentPosition + ":" + v
					rookMoves = append(rookMoves, s)
				}
				if len(rookMoves) > 0 {
					allMoves = append(allMoves, rookMoves)
				}
			}
			if board.Board[i][j].Name == constants.BLACK_PAWN {
				pawnOptions = util.GetPawnPaths(board.Board[i][j], board)
				pawnMoves := make([]string, 0)
				for _, v := range pawnOptions {
					s := "pawn:" + board.Board[i][j].CurrentPosition + ":" + v
					pawnMoves = append(pawnMoves, s)
				}
				if len(pawnMoves) > 0 {
					allMoves = append(allMoves, pawnMoves)
				}
			}
		}
	}
	return allMoves, nil
}

func chooseRandomMoveForRandomPiece(allMoves [][]string) (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	max := len(allMoves)
	pieceR := rand.Intn(max)

	switch pieceR {
	case 0:
		max := len(allMoves[0])
		nr := rand.Intn(max)
		return allMoves[0][nr], nil
	case 1:
		max := len(allMoves[1])
		nr := rand.Intn(max)
		return allMoves[1][nr], nil
	case 2:
		max := len(allMoves[2])
		nr := rand.Intn(max)
		return allMoves[2][nr], nil
	case 3:
		max := len(allMoves[3])
		nr := rand.Intn(max)
		return allMoves[3][nr], nil
	case 4:
		max := len(allMoves[4])
		nr := rand.Intn(max)
		return allMoves[4][nr], nil
	case 5:
		max := len(allMoves[5])
		nr := rand.Intn(max)
		return allMoves[5][nr], nil
	}
	return "-99", errors.New("There was a problem generating a random move for the computer")
}
