package computer

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/board"
	"go_chess/game/util"
	"strings"
	"testing"
)

func TestGenerateMoves(t *testing.T) {

	board := board.Game{}.InitializeBoard()

	res, _ := generateMoves(board)
	total := 0

	for _, v := range res {
		total += len(v)
	}

	assert.Equal(t, 24, total)
	assert.True(t, strings.Contains(res[3][0], "knight"))
	assert.True(t, strings.Contains(res[5][0], "pawn"))
}

//TODO: Uncomment this and complete this test
//func TestGenerateMoves2(t *testing.T) {
//
//	board := board.Game{}.InitializeBoard()
//
//	move1 := ""
//	move2 := ""
//
//	p := util.PlayMove("c7", board)
//	board, p = util.PlacePiece(p, "c5", board)
//
//	p2 := util.GetPieceFromPosition("b8", board)
//	board, _ = util.PlacePiece(p2, "c6", board)
//
//	board.PrintBoard()
//
//	res, _ := generateMoves(board)
//	total := 0
//
//	for _, v := range res {
//		total += len(v)
//	}
//
//	assert.Equal(t, 24, total)
//	assert.True(t, strings.Contains(res[3][0], "knight"))
//	assert.True(t, strings.Contains(res[5][0], "pawn"))
//}

func TestChooseRandomMoveForRandomPiece(t *testing.T) {
	board := board.Game{}.InitializeBoard()
	p := util.GetPieceFromPosition("c7", board)
	board, p = util.PlacePiece(p, "c6", board)
	options, _ := generateMoves(board)

	p = util.GetPieceFromPosition("f7", board)
	board, p = util.PlacePiece(p, "f6", board)

	options, _ = generateMoves(board)
	choice, _ := chooseRandomMoveForRandomPiece(options)
	assert.True(t, choice != "-99")
	parts := strings.Split(choice, ":")
	assert.Equal(t, 3, len(parts))
}
