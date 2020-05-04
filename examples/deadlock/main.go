package main

import (
	"fmt"
	"sync"
)

var data = []string{"Earth", "Mars", "Venus", "Saturn"}

// Based on "Understanding real-worldconcurrency bugs in go"
// by T. Tu, X. Liu, L. Song and Y. Zhang
func main() {
	var group sync.WaitGroup
	group.Add(len(data))
	for _, d := range data {
		go func(d string) {
			fmt.Printf("Processing: %s\n", d)
			defer group.Done()
		}(d)
		group.Wait() // This is causing the deadlock
	}
	// group.Wait() // Move here to fix it
}
