package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GetRating(wg *sync.WaitGroup) float64 {

	defer wg.Done()

	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Millisecond)
	rating := rand.Intn(5)
	return float64(rating)
}

func main() {
	var wg sync.WaitGroup
	var sum float64 = 0
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			sum += GetRating(&wg)
		}()
	}
	wg.Wait()

	averageRating := sum / 200

	fmt.Println("Average rating is: ", averageRating)
}
