package internal

import (
	"math"
)

type Point struct {
	X, Y float64
}

func ClosestPoint(start *Point, targets []Point) Point {
	if len(targets) == 0 {
		return *start
	}

	closest := targets[0]
	minDistance := manhattanDistance(*start, closest)

	for _, target := range targets[1:] {
		d := manhattanDistance(*start, target)
		if d < minDistance {
			closest = target
			minDistance = d
		}
	}

	return closest
}

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(a, b Point) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}
