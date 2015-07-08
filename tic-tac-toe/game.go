package main

import (
	"errors"
	"fmt"
)

type Game struct {
	Player1     Box
	Player2     Box
	CurrentTurn Box
	Grid        *Grid
	won         bool
}

var ErrInvalidPick = errors.New("invalid pick: O or X")
var ErrIllegalMove = errors.New("illegal move")
var ErrNotYourTurn = errors.New("not your turn yet")
var ErrGameWon = errors.New("game won")

func NewGame(p1 Box) (*Game, error) {
	if p1 == Blank {
		return nil, ErrInvalidPick
	}

	return &Game{
		Player1:     p1,
		Player2:     Box(-int(p1)),
		CurrentTurn: p1,
		Grid:        NewGrid(),
		won:         false,
	}, nil
}

func (g Game) String() string {
	template := "Game{CurrentTurn: %s, Player1: %s, Player2: %s\n%s}"
	return fmt.Sprintf(template, g.CurrentTurn.String(), g.Player1.String(), g.Player2.String(), g.Grid)
}

func (g *Game) PlaceMove(b Box, x, y int) error {
	if g.won {
		return ErrGameWon
	}
	if b != g.CurrentTurn {
		return ErrNotYourTurn
	}

	if !g.Grid.CanPlaceAMove(x, y) {
		return ErrIllegalMove
	}

	g.Grid[x][y] = g.CurrentTurn
	if g.CurrentTurn == g.Grid.CheckForWinner() {
		g.won = true
	}
	g.CurrentTurn = g.CurrentTurn.Alternate()

	return nil
}
