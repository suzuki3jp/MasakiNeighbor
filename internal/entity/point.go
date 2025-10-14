package entity

import "math"

type Point struct {
	X   float64
	Y   float64
	IsA bool
}

func SplitPointsByType(points []Point) (aPoints []Point, bPoints []Point) {
	for _, p := range points {
		if p.IsA {
			aPoints = append(aPoints, p)
		} else {
			bPoints = append(bPoints, p)
		}
	}
	return
}

func Distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}
