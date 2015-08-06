package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter = 0
var mutex = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func main() {
	// Set up the Seed for rand.Intn
	rand.Seed(time.Now().Unix())

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("Counter %d: %d\n", i, counter)
			mutex.Lock()
			counter += 1
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
}
