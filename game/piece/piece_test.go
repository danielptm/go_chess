package piece

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/constants"
	"testing"
)

func TestPiece_GetSuit(t *testing.T) {
	p := Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	res := p.GetSuit()
	assert.Equal(t, "WHITE", res)
}
