package computer

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/board"
	"strings"
	"testing"
)

func TestComputerDecides(t *testing.T) {

	board := board.Game{}.InitializeBoard()

	res, _ := generateMoves(board)
	total := 0

	for _, v := range res {
		total += len(v)
	}

	assert.Equal(t, 24, total)
	assert.True(t, strings.Contains(res[0][0], "knight"))
	assert.True(t, strings.Contains(res[1][0], "knight"))
	assert.True(t, strings.Contains(res[2][0], "pawn"))
}

func testChooseRandomMoveForRandomPiece(t *testing.T) {
	board := board.Game{}.InitializeBoard()

	responses := make([]string, 0)

	for i := 0; i < 100; i++ {
		res, _ := generateMoves(board)
		choice, _ := chooseRandomMoveForRandomPiece(res)
		responses = append(responses, choice)
	}

	first := responses[0]

	areDifferent := false

	for _, v := range responses {
		if v == first {
			areDifferent = false
		} else {
			areDifferent = true
			break
		}
	}

	assert.True(t, areDifferent)
}
