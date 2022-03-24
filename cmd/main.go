package main

import (
	"fmt"
	"go_chess/game"
)

func main() {
	fmt.Println("")
	fmt.Println("***")
	fmt.Println("Hey! Welcome to my chess game written in GO")
	fmt.Println("***")
	fmt.Println("")

	game := new(game.Game).InitializeBoard()
	game.PrintBoard()

}
