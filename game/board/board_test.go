package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializeBoard(t *testing.T) {
	g := Game{}.InitializeBoard()
	assert.Equal(t, 1, g.Board[0][1].CurrentX)
}
