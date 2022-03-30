package main

import (
	"bufio"
	"fmt"
	"go_chess/game/board"
	"os"
)

func main() {
	fmt.Println("")
	fmt.Println("***")
	fmt.Println("Hey! Welcome to my chess game written in GO")
	fmt.Println("***")
	fmt.Println("")

	game := new(board.Game).InitializeBoard()
	Run(game)

}

func Run(game board.Game) {
	println("")
	println("***")
	println("Please enter your moves in the format: <piece>:<start position>:<end position>")
	println("example: pawn:a2:a4")
	print("Press the 'enter' key to make the move")
	println("***")
	println("")
	winnerFound := false
	for !winnerFound {
		game.PrintBoard()
		println("")
		println("Your move:")
		println("")
		PlayerMoves()
		ComputerMoves()
	}
}

func PlayerMoves() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadString('\n')
	return res, err
}

func ComputerMoves() string {
	move := "move"
	println("")
	fmt.Printf("Computer'spa move: %s", move)
	println("")
	return "move"
}
