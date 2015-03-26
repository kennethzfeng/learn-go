package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter = 0
var mutex = &sync.Mutex{}
var done = make(chan bool)

func main() {
	// Set up the Seed for rand.Intn
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("Counter %d: %d\n", i, counter)
			mutex.Lock()
			counter += 1
			mutex.Unlock()

			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
