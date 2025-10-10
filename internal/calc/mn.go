package calc

import (
	"math"

	"github.com/suzuki3jp/mn/internal/entity"
)

type MnResult struct {
	R  float64
	Ro float64
	Re float64
	Z  float64
}

func Mn(points []entity.Point, A float64) MnResult {
	aPoints, bPoints := entity.SplitPointsByType(points)
	Na := float64(len(aPoints))
	Nb := float64(len(bPoints))
	N := Na + Nb
	na := Na / N
	nb := Nb / N

	sigmaDab := 0.0
	for _, pa := range aPoints {
		d := _getClosestPoint(pa, bPoints)
		sigmaDab += d
	}
	sigmaDba := 0.0
	for _, pb := range bPoints {
		d := _getClosestPoint(pb, aPoints)
		sigmaDba += d
	}

	ro := (sigmaDab + sigmaDba) / N
	re := na/(2*math.Sqrt(Nb/A)) + nb/(2*math.Sqrt(Na/A))
	R := ro / re
	sigmare := math.Sqrt((((na * (Na / A) * (4 - na*math.Pi)) + (nb * (Nb / A) * (4 - nb*math.Pi)) - (2 * na * nb * math.Pi * math.Sqrt((Na/A)*(Nb/A)))) / (4 * (Na / A) * (Nb / A) * math.Pi)) / N)
	Z := (ro - re) / sigmare

	return MnResult{
		R:  R,
		Ro: ro,
		Re: re,
		Z:  Z,
	}
}

func _getClosestPoint(point entity.Point, points []entity.Point) (distance float64) {
	if len(points) == 0 {
		panic("points is empty")
	}

	closest := points[0]
	distance = entity.Distance(point, closest)
	for _, p := range points[1:] {
		d := entity.Distance(point, p)
		if d < distance {
			closest = p
			distance = d
		}
	}
	return
}
