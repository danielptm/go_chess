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
	println("You can also type showcache to see removed pieces or showdetails to see a text description of the board or showboard to reprint the board")
	print("Press the 'enter' key to make the move")
	println("***")
	println("")
	winnerFound := false
	for !winnerFound {
		game.PrintBoard()
		game = PlayerMoves(game)
		game = ComputerMoves(game)
	}
}

func PlayerMoves(b board.Game) board.Game {
	isValid := false
	for !isValid {
		reader := bufio.NewReader(os.Stdin)
		if !b.HumanIsInCheck {
			println("")
			println("Your move:")
			println("")
		} else {
			println("")
			println("You're in check and it's your move:")
			println("")
		}
		res, _ := reader.ReadString('\n')
		res = res[0 : len(res)-1]
		if res == "showcache" {
			b.PrintCache()
			continue
		}
		if res == "showdetail" {
			b.PrintBoardDetail()
			continue
		}
		if res == "showboard" {
			b.PrintBoard()
			continue
		}
		hIsValid, err := util.CheckIfHumanMoveIsValid(res, b)
		if !hIsValid || err != nil {
			print("Your move was invalid.")
			isValid = false
		} else {
			isValid = hIsValid
			b = util.PlayMove(res, b)
		}
	}
	return b
}

func ComputerMoves(b board.Game) board.Game {
	move := computer.ComputerRandomlyDecides(b)
	println("Computer makes move: " + move)
	b = util.PlayMove(move, b)
	return b
}
