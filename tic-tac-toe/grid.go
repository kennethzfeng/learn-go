package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// Blank represents a [ ] box.
	Blank = 0
	// Circle represents a [O] box.
	Circle = -1
	// Cross represents a [X} box.
	Cross = 1
)

// Box is a unit [ ] in the grid
type Box int

var ErrUnableToAlternate = errors.New("unable to alternate")

const GridSize = 3

func (b Box) String() string {
	switch b {
	case Blank:
		return " "
	case Circle:
		return "O"
	case Cross:
		return "X"
	}

	return ""
}

func (b Box) Alternate() Box {
	if b == Circle {
		return Cross
	}
	if b == Cross {
		return Circle
	}
	return Blank
}

// Coordinate represents the location of a box in the grid
// coordinate is zero-based index starting from the top left corner
//
// For example, (0, 0) would be the box in the upper left hand corner
type Coordinate struct {
	X int
	Y int
}

// Grid is a 3-by-3 grid.
type Grid [GridSize][GridSize]Box

func (g Grid) Size() int {
	return GridSize
}

func (g Grid) boundCheck(c Coordinate) error {
	if c.X > g.Size()-1 || c.X < 0 {
		return errors.New("Coordinate out of bound")
	}
	if c.Y > g.Size()-1 || c.Y < 0 {
		return errors.New("Coordinate out of bound")
	}
	return nil
}

// BoxAtCoor returns the Box value at the coordinate
func (g Grid) BoxAtCoor(c Coordinate) (Box, error) {
	if err := g.boundCheck(c); err != nil {
		return Blank, err
	}
	return g[c.X][c.Y], nil
}

func (g *Grid) SetBoxAtCoor(b Box, c Coordinate) error {
	g[c.X][c.Y] = b
	return nil
}

// NewGrid return a new Grid representing the 3 by 3 Tic Tac Toe Grid
func NewGrid() *Grid {
	var g Grid

	for i, inner := range g {
		for j := range inner {
			g[i][j] = Blank
		}
	}

	return &g
}

func (g Grid) CanPlaceAMove(x, y int) bool {
	if x > g.Size()-1 || x < 0 {
		return false
	}
	if y > g.Size()-1 || y < 0 {
		return false
	}

	if g[x][y] != Blank {
		return false
	}

	return true
}

func (g Grid) String() string {
	var lines []string
	for i := range g {
		var items []string
		for j := range g[i] {
			items = append(items, g[i][j].String())
		}
		line := strings.Join(items, "|")
		// surround the line to start and end
		line = fmt.Sprintf("|%s|", line)
		lines = append(lines, line)
	}
	width := len(lines[0])
	heading := ""
	for i := 0; i < width; i++ {
		if i == 0 || i == width-1 {
			heading += "+"
		} else {
			heading += "-"
		}
	}

	var newLines []string

	newLines = append(newLines, heading)
	newLines = append(newLines, lines...)
	newLines = append(newLines, heading)

	return strings.Join(newLines, "\n")
}

// CheckForWinner checks whethere there is a winner
func (g Grid) CheckForWinner() Box {
	// Horizontial
	for _, inner := range g {
		sum := 0
		for _, v := range inner {
			sum += int(v)
		}
		if sum == Cross*g.Size() {
			return Cross
		}
		if sum == Circle*g.Size() {
			return Circle
		}
	}

	// Vertical
	for r := range g {
		sum := 0
		for c := range g[r] {
			b, _ := g.BoxAtCoor(Coordinate{X: c, Y: r})
			sum += int(b)
		}
		if sum == Cross*g.Size() {
			return Cross
		}
		if sum == Circle*g.Size() {
			return Circle
		}
	}

	// Diagonal
	// (0,0) (0,1) (0,2)
	// (1,0) (1,1) (1,2)
	// (2,0) (2,1) (2,2)
	sum := 0
	for i := range g {
		sum += int(g[i][i])
	}
	if sum == Cross*g.Size() {
		return Cross
	}
	if sum == Circle*g.Size() {
		return Circle
	}

	sum = 0
	for i := range g {
		sum += int(g[i][2-i])
	}
	if sum == Cross*g.Size() {
		return Cross
	}
	if sum == Circle*g.Size() {
		return Circle
	}

	return Blank
}
