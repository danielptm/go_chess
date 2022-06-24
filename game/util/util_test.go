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
	board := board.Game{}.InitializeBoard()
	p := GetPieceFromPosition("c7", board)
	paths := GetPawnPaths(p, board)
	assert.Equal(t, 2, len(paths))
	assert.Equal(t, "c5", paths[0])
	assert.Equal(t, "c6", paths[1])
}

func TestGetPawnPaths2(t *testing.T) {
	board := board.Game{}.InitializeBoard()
	p := GetPieceFromPosition("c7", board)
	p.HasMoved = true
	paths := GetPawnPaths(p, board)
	assert.Equal(t, 1, len(paths))
	assert.Equal(t, "c6", paths[0])
}

func TestGetRookPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_ROOK,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	b := board.Game{}.InitializeEmptyBoard()
	b, p = PlacePiece(p, "d4", b)
	paths := GetRookPaths(p, b)
	assert.Equal(t, 14, len(paths))
	assert.True(t, Contains("a4", paths))
}

func TestGetKingPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KING,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	b := board.Game{}.InitializeEmptyBoard()
	b, p = PlacePiece(p, "d4", b)
	paths := GetKingPaths(p, b)
	assert.Equal(t, 8, len(paths))
}

func TestGetQueenPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	np := GetCoordinates(p)
	g := board.Game{}.InitializeEmptyBoard()
	paths := GetQueenPaths(np, g)
	assert.Equal(t, 27, len(paths))
}

func TestGetBishopPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	g := board.Game{}.InitializeEmptyBoard()

	np := GetCoordinates(p)
	paths := GetBishopPaths(np, g)
	assert.Equal(t, 13, len(paths))
}

func TestGetDirectUpPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_PAWN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "d7", b)
	paths := GetDirectUpPaths(p, b)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("d7", paths))
}

func TestGetDirectUpPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	g := board.Game{}.InitializeEmptyBoard()
	np := GetCoordinates(p)
	paths := GetDirectUpPaths(np, g)
	assert.Equal(t, 4, len(paths))
	assert.True(t, Contains("d8", paths))
}

func TestGetDirectUpPaths3(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_PAWN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "d7", b)
	paths := GetDirectUpPaths(p, b)
	assert.Equal(t, 2, len(paths))
}

func TestGetUpRightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	g := board.Game{}.InitializeEmptyBoard()

	np := GetCoordinates(p)
	paths := GetUpRightPaths(np, g)
	assert.Equal(t, 4, len(paths))
	assert.True(t, Contains("g7", paths))
}

func TestGetUpRightPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	p2 := piece.Piece{
		Name:            constants.BLACK_PAWN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "g7", b)
	paths := GetUpRightPaths(p, b)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("g7", paths))
}

func TestGetUpRightPaths3(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	p2 := piece.Piece{
		Name:            constants.WHITE_PAWN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "g7", b)
	paths := GetUpRightPaths(p, b)
	assert.Equal(t, 2, len(paths))
	assert.True(t, Contains("f6", paths))
}

func TestGetRightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "g4", b)
	paths := GetRightPaths(p, b)
	assert.Equal(t, 3, len(paths))
}

func TestGetRightPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "g4", b)
	paths := GetRightPaths(p, b)
	assert.Equal(t, 2, len(paths))
}

func TestGetDownRightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d6", board)
	b, p2 = PlacePiece(p2, "g3", b)
	paths := GetDownRightPaths(p, b)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("e5", paths))
	assert.True(t, Contains("f4", paths))
	assert.True(t, Contains("g3", paths))
}

func TestGetDownRightPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d6", board)
	b, p2 = PlacePiece(p2, "g3", b)
	paths := GetDownRightPaths(p, b)
	assert.Equal(t, 2, len(paths))
}

func TestGetDownPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d6", board)
	b, p2 = PlacePiece(p2, "d3", b)
	np := GetCoordinates(p)
	paths := GetDownPaths(np, b)
	assert.Equal(t, 3, len(paths))
	assert.True(t, Contains("d5", paths))
}

func TestGetDownPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d6", board)
	b, p2 = PlacePiece(p2, "d3", b)
	np := GetCoordinates(p)
	paths := GetDownPaths(np, b)
	assert.Equal(t, 2, len(paths))
}

func TestGetDownLeftPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "a1", b)
	np := GetCoordinates(p)
	paths := GetDownLeftPaths(np, b)
	assert.Equal(t, 3, len(paths))
}

func TestGetDownLeftPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "a1", b)
	np := GetCoordinates(p)
	paths := GetDownLeftPaths(np, b)
	assert.Equal(t, 2, len(paths))
}

func TestGetLeftPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "a4", b)
	np := GetCoordinates(p)
	paths := GetLeftPaths(np, b)
	assert.Equal(t, 3, len(paths))
}

func TestGetLeftPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "a4", b)
	np := GetCoordinates(p)
	paths := GetLeftPaths(np, b)
	assert.Equal(t, 2, len(paths))
}

func TestGetUpLeftPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.BLACK_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "a7", b)
	np := GetCoordinates(p)
	paths := GetUpLeftPaths(np, b)
	assert.Equal(t, 3, len(paths))
}

func TestGetUpLeftPaths2(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	p2 := piece.Piece{
		Name:            constants.WHITE_BISHOP,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	b, p := PlacePiece(p, "d4", board)
	b, p2 = PlacePiece(p2, "a7", b)
	np := GetCoordinates(p)
	paths := GetUpLeftPaths(np, b)
	assert.Equal(t, 2, len(paths))
}

func TestUpBigLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := UpBigLeftL(p, board)
	assert.Equal(t, "c6", s[0])
}

func TestUpBigRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := UpBigRightL(p, board)
	assert.Equal(t, "e6", s[0])
}

func TestUpSmallLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := UpSmallLeftL(p, board)
	assert.Equal(t, "b5", s[0])
}

func TestUpSmallRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := UpSmallRightL(p, board)
	assert.Equal(t, "b5", s[0])
}

func TestRightBigUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := RightBigUpL(p, board)
	assert.Equal(t, "e6", s[0])
}

func TestRightBigDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := RightBigDownL(p, board)
	assert.Equal(t, "e2", s[0])
}

func TestRightSmallUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := RightSmallUpL(p, board)
	assert.Equal(t, "f5", s[0])
}

func TestRightSmallDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := RightSmallDownL(p, board)
	assert.Equal(t, "f3", s[0])
}

func TestDownBigLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := DownBigLeftL(p, board)
	assert.Equal(t, "c2", s[0])
}

func TestDownBigRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := DownBigRightL(p, board)
	assert.Equal(t, "e2", s[0])
}

func TestDownBigRightL2(t *testing.T) {
	board := board.Game{}.InitializeBoard()
	p := GetPieceFromPosition("b8", board)
	s := DownBigRightL(p, board)
	np := GetPieceFromPosition("b8", board)
	s2 := DownBigLeftL(np, board)
	assert.Equal(t, "c6", s[0])
	assert.Equal(t, "a6", s2[0])
}

func TestDownSmallLeftL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := DownSmallLeftL(p, board)
	assert.Equal(t, "b3", s[0])
}

func TestDownSmallRightL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := DownSmallRightL(p, board)
	assert.Equal(t, "f3", s[0])
}

func TestLeftBigUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := LeftBigUpL(p, board)
	assert.Equal(t, "b5", s[0])
}

func TestLeftBigDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := LeftBigDownL(p, board)
	assert.Equal(t, "b3", s[0])
}

func TestLeftSmallUpL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := LeftSmallUpL(p, board)
	assert.Equal(t, "c6", s[0])
}

func TestLeftSmallDownL(t *testing.T) {
	p := piece.Piece{
		Name:            constants.WHITE_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	s := LeftSmallDownL(p, board)
	assert.Equal(t, "c2", s[0])
}

func TestGetKnightPaths(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "d4", board)
	paths := GetKnightPaths(p, board)
	assert.Equal(t, 16, len(paths))
}

func TestGetKnightPaths2(t *testing.T) {
	board := board.Game{}.InitializeBoard()
	p := GetPieceFromPosition("b8", board)
	moves := GetKnightPaths(p, board)
	assert.Equal(t, 4, len(moves))
	assert.True(t, Contains("a6", moves))
	assert.True(t, Contains("c6", moves))
}

func TestGetKnightPathsWithSomeInvalid(t *testing.T) {
	p := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "g6",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, p = PlacePiece(p, "g6", board)
	paths := GetKnightPaths(p, board)
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
	newBoard, p := PlacePiece(p, "a6", b)
	assert.Equal(t, newBoard.Board[2][0].Name, constants.BLACK_PAWN)
}

func TestIsSameColor(t *testing.T) {
	p1 := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	p2 := piece.Piece{
		Name:            constants.WHITE_PAWN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	res := IsSameColor(p1.Name, p2.Name)

	assert.Equal(t, false, res)
}

func TestIsSameColor2(t *testing.T) {
	p1 := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	p2 := piece.Piece{
		Name:            constants.BLACK_QUEEN,
		CurrentPosition: "d4",
		HasMoved:        true,
	}

	res := IsSameColor(p1.Name, p2.Name)

	assert.Equal(t, true, res)
}

func TestGetPieceFromPosition(t *testing.T) {
	p1 := piece.Piece{
		Name:            constants.BLACK_KNIGHT,
		CurrentPosition: "d4",
		HasMoved:        true,
	}
	board := board.Game{}.InitializeEmptyBoard()
	board, _ = PlacePiece(p1, "d4", board)
	np := GetPieceFromPosition("d4", board)

	assert.Equal(t, constants.BLACK_KNIGHT, np.Name)
}

func TestPlayMove(t *testing.T) {
	b := board.Game{}.InitializeBoard()
	move := "pawn:c7:c6"
	b = PlayMove(move, b)
	newc7 := GetPieceFromPosition("c7", b)
	newc6 := GetPieceFromPosition("c6", b)
	assert.Equal(t, "", newc7.Name)
	assert.Equal(t, constants.BLACK_PAWN, newc6.Name)
}

func TestTakePiece(t *testing.T) {
	b := board.Game{}.InitializeEmptyBoard()
	p := piece.Piece{Name: constants.BLACK_BISHOP, CurrentPosition: "b7", HasMoved: true}
	p2 := piece.Piece{Name: constants.WHITE_PAWN, CurrentPosition: "d5", HasMoved: true}
	b, _ = PlacePiece(p, "b7", b)
	b, _ = PlacePiece(p2, "d5", b)
	b = TakePiece("d5", p, b)
	res := GetPieceFromPosition("d5", b)
	res2 := GetPieceFromPosition("b7", b)
	assert.Equal(t, p.Name, res.Name)
	assert.Equal(t, "", res2.Name)
}
