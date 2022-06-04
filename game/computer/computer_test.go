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
