package board

import "testing"

func TestGame_PrintBoard(t *testing.T) {
	x := [2][2]string{{"a", "b"}, {"c", "d"}}

	println(x[0][0])
	println(x[0][1])
	println(x[1][0])
	println(x[1][1])
}
