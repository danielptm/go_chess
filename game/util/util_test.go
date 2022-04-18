package util

import (
	"github.com/stretchr/testify/assert"
	"go_chess/game/board"
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

func TestDirectRightForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := DirectRightForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "f4", res.CurrentPosition)
}

func TestDirectLeftForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := DirectLeftForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "b4", res.CurrentPosition)
}

func TestUpRightForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := UpRightForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "f6", res.CurrentPosition)
}

func TestDownRightForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := DownRightForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "f2", res.CurrentPosition)
}

func TestUpLeftForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := UpLeftForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "b6", res.CurrentPosition)
}

func TestDownLeftForSpaces(t *testing.T) {
	p := piece.Piece{
		Name:     constants.BLACK_BISHOP,
		CurrentX: 3,
		CurrentY: 4,
	}

	np := GetBoardPosition(p)

	np2 := GetCoordinates(np)

	resY, resX := DownLeftForSpaces(np2.CurrentX, np2.CurrentY, 2)

	res := GetBoardPosition(piece.Piece{CurrentX: resX, CurrentY: resY})

	assert.Equal(t, "b2", res.CurrentPosition)
}

func TestGetPawnPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_PAWN,
		CurrentPosition: "d2",
	}
	np := GetCoordinates(p)
	paths := GetPawnPaths(np)
	assert.Equal(t, 2, len(paths))
	assert.Equal(t, "d4", paths[0])
	assert.Equal(t, "d3", paths[1])
}

func TestGetPawnPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_PAWN,
		CurrentPosition: "d2",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetPawnPaths(np)
	assert.Equal(t, 1, len(paths))
	assert.Equal(t, "d3", paths[0])
}

func TestGetRookPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_ROOK,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetRookPaths(np)
	assert.Equal(t, 14, len(paths))
	assert.True(t, Contains("a4", paths))
}

func TestGetKingPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetKingPaths(np)
	assert.Equal(t, 8, len(paths))
}

func TestGetQueenPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetQueenPaths(np)
	assert.Equal(t, 27, len(paths))
}

func TestGetBishopPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetBishopPaths(np)
	assert.Equal(t, 13, len(paths))
}

func TestGetDirectUpPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetDirectUpPaths(np)
	assert.Equal(t, 4, len(paths))
	assert.True(t, Contains("d8", paths))
}

func TestGetUpRightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetUpRightPaths(np)
	assert.Equal(t, 4, len(paths))
	assert.True(t, Contains("g7", paths))

}

func TestGetRightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetRightPaths(np)
	assert.Equal(t, 4, len(paths))
	assert.True(t, Contains("h4", paths))
}

func TestGetDownRightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetDownRightPaths(np)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("g1", paths))

}

func TestGetDownPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetDownPaths(np)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("d1", paths))
}

func TestGetDownLeftPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetDownLeftPaths(np)
	assert.Equal(t, 3, len(paths))
}

func TestGetLeftPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetLeftPaths(np)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("a4", paths))
}

func TestGetUpLeftPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetUpLeftPaths(np)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("a7", paths))
}

func TestUpBigLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := UpBigLeftL(np)
	assert.Equal(t, "c6", s[0])
}

func TestUpBigRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := UpBigRightL(np)
	assert.Equal(t, "e6", s[0])
}

func TestUpSmallLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := UpSmallLeftL(np)
	assert.Equal(t, "b5", s[0])
}

func TestUpSmallRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := UpSmallRightL(np)
	assert.Equal(t, "b5", s[0])
}

func TestRightBigUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := RightBigUpL(np)
	assert.Equal(t, "e6", s[0])
}

func TestRightBigDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := RightBigDownL(np)
	assert.Equal(t, "e2", s[0])
}

func TestRightSmallUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := RightSmallUpL(np)
	assert.Equal(t, "f5", s[0])
}

func TestRightSmallDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := RightSmallDownL(np)
	assert.Equal(t, "f3", s[0])
}

func TestDownBigLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := DownBigLeftL(np)
	assert.Equal(t, "c2", s[0])
}

func TestDownBigRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := DownBigRightL(np)
	assert.Equal(t, "e2", s[0])
}

func TestDownSmallLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := DownSmallLeftL(np)
	assert.Equal(t, "b3", s[0])
}

func TestDownSmallRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := DownSmallRightL(np)
	assert.Equal(t, "f3", s[0])
}

func TestLeftBigUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := LeftBigUpL(np)
	assert.Equal(t, "b5", s[0])
}

func TestLeftBigDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := LeftBigDownL(np)
	assert.Equal(t, "b3", s[0])
}

func TestLeftSmallUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := LeftSmallUpL(np)
	assert.Equal(t, "c6", s[0])
}

func TestLeftSmallDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	s := LeftSmallDownL(np)
	assert.Equal(t, "c2", s[0])
}

func TestGetKnightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetKnightPaths(np)
	assert.Equal(t, 16, len(paths))
}

func TestGetKnightPathsWithSomeInvalid(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "g6",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	paths := GetKnightPaths(np)
	assert.Equal(t, 13, len(paths))
}

func TestIsInBounds(t *testing.T) {
	x := 8
	y := 4
	res := IsInbounds(x, y)
	assert.Equal(t, false, res)
}

func TestIsInBounds2(t *testing.T) {
	x := 6
	y := 4
	res := IsInbounds(x, y)
	assert.Equal(t, true, res)
}

func TestIsInBounds3(t *testing.T) {
	x := -1
	y := 4
	res := IsInbounds(x, y)
	assert.Equal(t, false, res)
}

func TestIsInBounds4(t *testing.T) {
	x := 9
	y := 4
	res := IsInbounds(x, y)
	assert.Equal(t, false, res)
}

func TestPlacePiece(t *testing.T) {
	g := board.Game{}
	p := piece.Piece{
		Name:     constants.BLACK_PAWN,
		HasMoved: true,
	}
	b := g.InitializeEmptyBoard()
	newBoard := PlacePiece(p, "a6", b)
	assert.Equal(t, newBoard.Board[2][0].Name, constants.BLACK_PAWN)
}
