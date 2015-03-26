package main

import (
	"fmt"
	"time"
)

func say(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(name)
	}
}
func main() {
	go say("Tim")
	say("Kenneth")
}
