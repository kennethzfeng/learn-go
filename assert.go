//
// Demo for using assert
//

package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func main() {
	fmt.Println("HellO")
}

func testMain(t *testing.T) {
	assert.Equal(t, 1, 1, "1 == 1")
}
