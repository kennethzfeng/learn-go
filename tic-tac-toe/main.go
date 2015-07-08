// Tic-Tac-Toe
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	g, err := NewGame(Cross)
	if err != nil {
		log.Fatalf("unable to create game: %s", err)
		os.Exit(1)
	}
	move(g, 0, 0)
	move(g, 1, 1)
	move(g, 0, 1)
	move(g, 1, 2)
	move(g, 0, 2)
	move(g, 2, 2)
	move(g, 2, 0)

}

func move(g *Game, x, y int) {
	err := g.PlaceMove(g.CurrentTurn, x, y)
	if err != nil {
		log.Fatalf("unable to make a move: %s", err)
		os.Exit(1)
	}
	fmt.Println(g.Grid)
	fmt.Println("Current Turn: " + g.CurrentTurn.String())
}
