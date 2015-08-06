package main

import (
	"fmt"
	"testing"
)

func TestCheckForWinnerWithHorizontial(t *testing.T) {
	g := NewGrid()

	g[0][0] = Cross
	g[0][1] = Cross
	g[0][2] = Cross

	winner := g.CheckForWinner()
	fmt.Println(g)

	if winner != Cross {
		t.Fatalf("Cross should be the winner")
	}
}

func TestCheckForWinnerWithVertical(t *testing.T) {
	g := NewGrid()

	g[0][0] = Circle
	g[1][0] = Circle
	g[2][0] = Circle

	winner := g.CheckForWinner()
	fmt.Println(g)

	if winner != Circle {
		t.Fatalf("Circle should be the winner")
	}
}

func TestCheckForWinnerWithDia1(t *testing.T) {
	g := NewGrid()

	g[0][0] = Circle
	g[1][1] = Circle
	g[2][2] = Circle

	winner := g.CheckForWinner()
	fmt.Println(g)

	if winner != Circle {
		t.Fatalf("Circle should be the winner")
	}
}

func TestCheckForWinnerWithDia2(t *testing.T) {
	g := NewGrid()

	g[0][2] = Circle
	g[1][1] = Circle
	g[2][0] = Circle

	winner := g.CheckForWinner()
	fmt.Println(g)

	if winner != Circle {
		t.Fatalf("Circle should be the winner")
	}
}
