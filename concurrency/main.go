package main

import (
	"fmt"
	"math/rand"
	"time"
)

var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Counter %d: %d\n", i, rand.Intn(5))
		}(i)
	}

	time.Sleep(10)
}
