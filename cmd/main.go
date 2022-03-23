package main

import (
	"fmt"
	game "go_chess/game"
)

func main() {
	fmt.Println("")
	fmt.Println("***")
	fmt.Println("Hey! Welcome to my chess game written in GO")
	fmt.Println("***")
	fmt.Println("")

	board := new(game.Game).InitializeBoard()

	PrintBoard(board)

}

func PrintBoard(game game.Game) {
	fmt.Println("*********** NEXT TURN ***********")
	for i := 0; i < 8; i++ {
		print("| ")
		print(game.Board[i][0])
		print(" | ")
		print(game.Board[i][1])
		print(" | ")
		print(game.Board[i][2])
		print(" | ")
		print(game.Board[i][3])
		print(" | ")
		print(game.Board[i][4])
		print(" | ")
		print(game.Board[i][5])
		print(" | ")
		print(game.Board[i][6])
		print(" | ")
		print(game.Board[i][7])
		print(" | ")
		println("")
	}
	println("")
}
