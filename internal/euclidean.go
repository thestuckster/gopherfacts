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
	minDistance := distance(*start, closest)

	for _, target := range targets[1:] {
		d := distance(*start, target)
		if d < minDistance {
			closest = target
			minDistance = d
		}
	}

	return closest
}

// distance calculates the Euclidean distance between two points
func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}
