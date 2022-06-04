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

	allMoves := [][]string{kingOptions, queenOptions, bishopOptions, knightOptions, rookOptions, pawnOptions}

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
					allMoves[0] = append(allMoves[0], kingMoves...)
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
					allMoves[1] = append(allMoves[1], queenMoves...)
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
					allMoves[2] = append(allMoves[2], bishopMoves...)
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
					allMoves[3] = append(allMoves[3], knightMoves...)
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
					allMoves[4] = append(allMoves[4], rookMoves...)
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
					allMoves[5] = append(allMoves[5], pawnMoves...)
				}
			}
		}
	}
	return allMoves, nil
}

//TODO: This fails sometimes because it will send in an array allMoves that has a length greater than 5
func chooseRandomMoveForRandomPiece(a [][]string) (string, error) {
	allMoves := make([][]string, 0)
	for _, v := range a {
		if len(v) > 0 {
			allMoves = append(allMoves, v)
		}
	}
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	max := len(allMoves)
	pieceR := rand.Intn(max - 1)

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
