package computer

import (
	"errors"
	"go_chess/game/board"
	"go_chess/game/util"
	"math/rand"
	"time"
)

func ComputerRandomlyDecides(board board.Game) string {
	res, _ := util.GenerateMoves(true, board)
	move, _ := chooseRandomMoveForRandomPiece(res)
	return move
}

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
