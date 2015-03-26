// walker walks throough the entire subtree
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s with %d bytes\n", path, f.Size())

	return nil
}

func main() {
	flag.Parse()
	path := flag.Arg(0)
	filepath.Walk(path, walkFunc)
}
