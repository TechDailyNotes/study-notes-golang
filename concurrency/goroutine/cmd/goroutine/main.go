package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Total time: ", time.Since(start))
	}()

	ninjas := []string {"Hattori", "Yoshi", "Kuma", "Fuma"}

	for _, ninja := range ninjas {
		wg.Add(1)
		go attack(ninja)
	}

	wg.Wait()
}

func attack(ninja string) {
	fmt.Println("Attack: ", ninja)
	time.Sleep(time.Second)
	wg.Done()
}