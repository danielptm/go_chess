package main

import (
	"bufio"
	"fmt"
	"go_chess/game/board"
	"go_chess/game/computer"
	"go_chess/game/util"
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
		game = PlayerMoves(game)
		game = ComputerMoves(game)
	}
}

func PlayerMoves(b board.Game) board.Game {
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	b = util.PlayMove(res, b)
	return b
}

func ComputerMoves(b board.Game) board.Game {
	move := computer.ComputerRandomlyDecides(b)
	println("Computer makes move: " + move)
	b = util.PlayMove(move, b)
	return b
}
