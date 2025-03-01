//With Goroutine channel
//-------------------------

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RoundResult stores the data for each round
type RoundResult struct {
	Round         int
	PeopleCrossed int
	TimeTaken     time.Duration
}

func walkRound(round int, ch chan RoundResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate number of people crossed (random 1 to 10)
	people := rand.Intn(10) + 1

	// Simulate time taken per round (random 1 to 5 seconds)
	duration := time.Duration(rand.Intn(5)+1) * time.Second
	time.Sleep(duration) // Simulate walking round

	ch <- RoundResult{Round: round, PeopleCrossed: people, TimeTaken: duration}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan RoundResult, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go walkRound(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalTime := time.Duration(0)
	totalPeople := 0

	fmt.Println("Round Summary:")
	for result := range ch {
		fmt.Printf("Round %d: Crossed %d people, Time Taken: %v\n", result.Round, result.PeopleCrossed, result.TimeTaken)
		totalPeople += result.PeopleCrossed
		totalTime += result.TimeTaken
	}

	fmt.Printf("\nTotal People Crossed: %d\n", totalPeople)
	fmt.Printf("Total Time Taken: %v\n", totalTime)
}
