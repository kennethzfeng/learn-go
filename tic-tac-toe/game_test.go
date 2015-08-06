package main

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	g, err := NewGame(Circle)
	if err != nil {
		t.Fatalf("it shouldn't have any error, but got %s", err)
	}
	if g.Player2 != Cross {
		t.Fatalf("Player2 should have Cross, but got %s", g.Player2.String())
	}
}

func TestPlaceAMoveChange(t *testing.T) {
	g, err := NewGame(Cross)
	if err != nil {
		t.Fatalf("unable to create game: %s", err)
	}
	fmt.Println(g)
	if g.CurrentTurn != Cross {
		t.Fatalf("the current move should X")
	}
	err = g.PlaceMove(Cross, 0, 0)
	if err != nil {
		t.Fatalf("unable to make a move: %s", err)
	}
	fmt.Println(g)
	if g.Grid[0][0] != Cross {
		t.Fatalf("the move at (0, 0) should X")
	}

	if g.CurrentTurn != Circle {
		t.Fatalf("the current move should O")
	}

}
