package point

import "math"

type Point struct {
	dim    int
	coord []float64
}

var NumOfCalls uint32 = 0

func EuclideanDistance(a, b Point) float64 {
	if a.dim != b.dim {
		panic("Both points must have same dimension")
	}
	NumOfCalls++
	sum := 0.
	for i := 0; i < a.dim; i++ {
		delta := a.GetCoord()[i] - b.GetCoord()[i]
		sum += delta * delta
	}
	return math.Sqrt(sum)
}

func CreatePoint(coords ...float64) Point {
	return Point{len(coords), coords}
}

func (p Point) GetCoord() []float64 {
	return p.coord
}

func (p Point) GetDimension() int {
	return p.dim
}
