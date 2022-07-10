package ai

import (
	"github.com/emirpasic/gods/stacks"
	"github.com/emirpasic/gods/stacks/arraystack"
	"go_chess/game/board"
	"go_chess/game/util"
)

type Dfs struct {
	stack stacks.Stack
}

func (d *Dfs) Init() {
	d.stack = arraystack.New()
}

func CreateRootVertices(board board.Game) []Vertex {
	pieceMoves, _ := util.GenerateMoves(true, board)
	vertices := make([]Vertex, 0)

	for _, pm := range pieceMoves {
		for _, m := range pm {
			v := Vertex{}
			v.Init(m, make([]Vertex, 0), board)
			vertices = append(vertices, v)
		}
	}
	return vertices
}

//func (d *Dfs) Dfs(vertices []Vertex, stopPoint int) []Vertex {
//
//	if stopPoint < 0 {
//		for i, _ := range vertices {
//			vertices[i] = d.iterate(vertices[i])
//		}
//	} else {
//		for i := 0; i < stopPoint; i++ {
//			vertices[i] = d.iterate(vertices[i])
//		}
//	}
//	return vertices
//}

//func (d *Dfs) iterate(v Vertex) Vertex {
//	v.Visited = true
//	d.stack.Push(v)
//	for !d.stack.Empty() {
//		curr, _ := d.stack.Pop()
//		moves := util.GenerateMoves(true, curr.)
//	}
//	return v
//}
