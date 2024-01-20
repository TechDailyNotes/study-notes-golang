package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Total time: ", time.Since(start))
	}()

	attacked := make(chan bool)

	ninja := "Hattori"
	go attack(ninja, attacked)
	fmt.Println(<-attacked)
}

func attack(ninja string, attacked chan bool) {
	fmt.Println("Attack: ", ninja)
	time.Sleep(time.Second)
	attacked <- true
}