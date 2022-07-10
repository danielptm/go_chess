package ai

import (
	"github.com/stretchr/testify/assert"
	board2 "go_chess/game/board"
	"testing"
)

func TestCreateRootVertices(t *testing.T) {
	board := board2.Game{}.InitializeSmallTestBoard()
	board.PrintBoard()
	res := CreateRootVertices(board)
	assert.Equal(t, 16, len(res))
}

//func TestDfs(t *testing.T) {
//	board := board2.Game{}.InitializeSmallTestBoard()
//	vertices := CreateRootVertices(board)
//	d := Dfs{}
//	d.Init()
//	d.Dfs(vertices, 3)
//}
