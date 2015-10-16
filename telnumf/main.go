// telnumf takes an input of a string of characters and turns it to numbers
// using telephone keypad
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) <= 0 {
		fmt.Println("no input provided.")
		os.Exit(1)
	}

	input := flag.Args()[0]
	for _, c := range strings.ToUpper(input) {
		switch c {
		case 'A', 'B', 'C':
			fmt.Printf("%d", 2)
		case 'D', 'E', 'F':
			fmt.Printf("%d", 3)
		case 'G', 'H', 'I':
			fmt.Printf("%d", 4)
		case 'J', 'K', 'L':
			fmt.Printf("%d", 5)
		case 'M', 'N', 'O':
			fmt.Printf("%d", 6)
		case 'P', 'Q', 'R', 'S':
			fmt.Printf("%d", 7)
		case 'T', 'U', 'V':
			fmt.Printf("%d", 8)
		case 'W', 'X', 'Y', 'Z':
			fmt.Printf("%d", 9)
		default:
		}
	}
	fmt.Print("\n")
}
