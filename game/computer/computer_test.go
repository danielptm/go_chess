package computer

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/board"
	"go_chess/game/constants"
	"go_chess/game/piece"
	"go_chess/game/util"
	"testing"
)

func TestComputerDecides(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = util.PlacePiece(p, "d4", board)

	res, _ := ComputerDecidesRandomly(board)

	assert.True(t, len(res) > 0)

}
