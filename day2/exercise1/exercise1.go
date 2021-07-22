package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mux sync.Mutex
	v   map[rune]int
}

// Inc increments the counter for the given key after locking the map.
func (c *SafeCounter) Inc(key rune) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func CalculateFrequency(s string, c *SafeCounter, wg *sync.WaitGroup) {

	defer wg.Done()

	for _, r := range s {
		go c.Inc(r)
	}
}

func ConcurrentFrequencyCounter(strings []string) SafeCounter {
	c := SafeCounter{v: make(map[rune]int)}

	var wg sync.WaitGroup

	// Start a goroutine to calculate the frequency map for each string.
	for _, s := range strings {
		wg.Add(1)
		go func(s string) {
			CalculateFrequency(s, &c, &wg)
		}(s)
	}

	wg.Wait()
	return c
}

func main() {
	strings := []string{"abac", "daef", "ffghid", "ekl"}

	freqMap := ConcurrentFrequencyCounter(strings)

	for k, v := range freqMap.v {
		fmt.Println(string(k), v)
	}
}
