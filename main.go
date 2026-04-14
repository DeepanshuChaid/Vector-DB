package main

import (
	"fmt"
	"math"
)

// Vector represents our embedding
type Vector []float32

// EuclideanDistance calculates L2 distance
func EuclideanDistance(a, b Vector) float64 {
	var sum float64
	for i := range a {
		diff := float64(a[i] - b[i])
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func main() {
	// Our "Database"
	db := map[string]Vector{
		"apple":  {0.1, 0.2, 0.3},
		"banana": {0.9, 0.8, 0.7},
		"cat":    {-0.1, 0.5, -0.2},
	}

	// Query vector (something similar to "apple")
	query := Vector{0.12, 0.18, 0.31}

	var bestMatch string
	minDist := math.MaxFloat64

	for label, vec := range db {
		dist := EuclideanDistance(query, vec)
		fmt.Printf("Distance to %s: %.4f\n", label, dist)
		if dist < minDist {
			minDist = dist
			bestMatch = label
		}
	}

	fmt.Printf("\nClosest match: %s\n", bestMatch)
}