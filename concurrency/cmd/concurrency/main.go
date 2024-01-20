package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Done in", time.Since(start))
	}()

	evilNinjas := []string{"Hank", "Kawasaki", "Hayabusa", "Hattori"}

	for _, ninja := range evilNinjas {
		attack(ninja)
	}
}

func attack(target string) {
	fmt.Println("Attacking", target)
	time.Sleep(time.Second)
}