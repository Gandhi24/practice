package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

func CalculateFrequency(s string, c *SafeCounter, wg *sync.WaitGroup) {

	defer wg.Done()

	for _, r := range s {
		go c.Inc(string(r))
	}
}

func ConcurrentFrequencyCounter(strings []string) SafeCounter {
	c := SafeCounter{v: make(map[string]int)}

	var wg sync.WaitGroup

	// Start a goroutine to calculate the frequency map for each string.
	for _, s := range strings {
		go func(s string) {
			wg.Add(1)
			CalculateFrequency(s, &c, &wg)
		}(s)
	}

	wg.Wait()
	return c
}

func main() {
	strings := []string{"abac", "daef", "ghi", "kl"}

	freqMap := ConcurrentFrequencyCounter(strings)

	fmt.Println(freqMap.v)
}
