package util

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/constants"
	"go_chess/game/piece"
	"testing"
)

func TestDirectUpForSpaces(t *testing.T) {
	piece := piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "a2"}
	x, y := piece.GetCoordinates()
	newY, newX := DirectUpForSpaces(x, y, 2)

	assert.Equal(t, 4, newY)
	assert.Equal(t, 0, newX)

}
