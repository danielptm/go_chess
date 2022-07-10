package ai

import "go_chess/game/board"

type Vertex struct {
	Move         string
	AdjacentList []Vertex
	ResultBoard  board.Game
	Visited      bool
}

func (v *Vertex) Init(move string, adjacentList []Vertex, resultBoard board.Game) {
	v.Move = move
	v.ResultBoard = resultBoard
	v.AdjacentList = adjacentList
}

func (v *Vertex) Add(newV Vertex) {
	v.AdjacentList = append(v.AdjacentList, newV)
}
