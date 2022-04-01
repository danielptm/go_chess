package piece

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/constants"
	"testing"
)

func TestPiece_GetArrayPosition(t *testing.T) {
	pos := "a8"
	p := Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: pos,
	}

	x, y := p.GetCoordinates()

	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)
}

func TestPiece_GetArrayPosition2(t *testing.T) {
	pos := "d4"
	p := Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: pos,
	}

	x, y := p.GetCoordinates()

	assert.Equal(t, 3, x)
	assert.Equal(t, 4, y)
}

func TestPiece_GetArrayPosition3(t *testing.T) {
	pos := "d4"
	p := Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: pos,
	}

	p.GetCoordinates()
	res2 := p.GetBoardPosition()

	assert.Equal(t, "d4", res2)
}
