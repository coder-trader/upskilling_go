package main

import (
	"fmt"
	"math/rand"
)

func myrand(start, end int) {
	rand.Seed(42) // Set a fixed seed for reproducibility
	randomNumber := rand.Intn(end-start+1) + start
	fmt.Printf("Random number between %d and %d: %d\n", start, end, randomNumber)
}
