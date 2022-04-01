package util

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/constants"
	"go_chess/game/piece"
	"testing"
)

func TestDirectUpForSpaces(t *testing.T) {
	piece := piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "a2"}
	np := GetCoordinates(piece)
	newY, newX := DirectUpForSpaces(np.CurrentX, np.CurrentY, 2)

	assert.Equal(t, 4, newY)
	assert.Equal(t, 0, newX)
}

func TestPiece_GetCoordinates(t *testing.T) {
	pos := "a8"
	p := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: pos,
	}

	np := GetCoordinates(p)

	assert.Equal(t, 0, np.CurrentX)
	assert.Equal(t, 0, np.CurrentY)
}

func TestPiece_GetCoordinates2(t *testing.T) {
	pos := "d4"
	p := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: pos,
	}

	np := GetCoordinates(p)

	assert.Equal(t, 3, np.CurrentX)
	assert.Equal(t, 4, np.CurrentY)
}

func TestPiece_GetPosition(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	res := GetBoardPosition(p)

	assert.Equal(t, "d4", res.CurrentPosition)
}

func TestDirectDownForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := DirectDownForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "d2", res.CurrentPosition)
}
