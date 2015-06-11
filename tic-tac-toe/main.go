// Tic-Tac-Toe
package main

const (
	Blank  = 0
	Circle = -1
	Cross  = 1
)

type Box int

// Grid is a 3-by-3 grid.
type Grid [3][3]Box

func NewGrid() *Grid {
	var g Grid

	for i, inner := range g {
		for j, _ := range inner {
			g[i][j] = Blank
		}
	}

	return &g
}

func (g Grid) CheckForWinner() int {
	// Horizontial
	for i, inner := range g {
		var suite int

		for j, v := range inner {
			if v == Blank {
				break
			} else {
				var suite
			}
		}
	}
}

func main() {
	g := NewGrid()
	println(g[1][1])

}
