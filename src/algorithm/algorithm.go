package algorithm

import (
	"math"
	"tucil/stima/pairit/point"
)

func QuickSort[T any](data []T, compareFunc func(T, T) int) {
	if len(data) <= 1 {
		return
	}
	// partisi
	pivotIdx := len(data) - 1
	pivot := data[pivotIdx]

	p := -1

	for q := 0; q < pivotIdx; q++ {
		if compareFunc(data[q], pivot) <= 0 {
			p++
			data[p], data[q] = data[q], data[p]
		}
	}

	data[p+1], data[pivotIdx] = data[pivotIdx], data[p+1]

	QuickSort(data[:p+1], compareFunc)
	QuickSort(data[p+1:], compareFunc)
}

func BruteForceFCP(points []point.Point) (*point.Point, *point.Point, float64) {
	p1 := &points[0]
	p2 := &points[1]
	min := point.EuclideanDistance(*p1, *p2)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := &points[i]
			b := &points[j]
			d := point.EuclideanDistance(*a, *b)
			if d < min {
				min = d
				p1 = a
				p2 = b
			}
		}
	}
	return p1, p2, min
}

func fcpImpl(sortedPoints []point.Point) (*point.Point, *point.Point, float64) {
	n := len(sortedPoints)

	if n == 2 {
		a := &sortedPoints[0]
		b := &sortedPoints[1]
		return a, b, point.EuclideanDistance(*a, *b)
	}

	mid := n/2 + 1
	if mid%2 == 1 {
		mid--
	}

	s1 := sortedPoints[:mid]
	var s2 []point.Point
	if n%2 == 1 {
		s2 = sortedPoints[mid-1:]
	} else {
		s2 = sortedPoints[mid:]
	}

	a1, b1, d1 := fcpImpl(s1)
	a2, b2, d2 := fcpImpl(s2)

	var (
		a, b *point.Point
		d    float64
	)

	if d1 < d2 {
		a, b, d = a1, b1, d1
	} else {
		a, b, d = a2, b2, d2
	}

	i := len(s1) - 1
	for i >= 0 && s1[i].GetCoord()[0] > sortedPoints[n/2].GetCoord()[0]-d {
		i--
	}

	j := 0
	for j < len(s2) && s2[j].GetCoord()[0] < sortedPoints[n/2].GetCoord()[0]+d {
		j++
	}

	var s []point.Point
	if n%2 == 1 {
		s = append(s1[i+1:], s2[1:j]...)
	} else {
		s = append(s1[i+1:], s2[:j]...)
	}

	for i = 0; i < len(s); i++ {
		for j = i + 1; j < len(s); j++ {
			a3 := &s[i]
			b3 := &s[j]
			ok := true
			for k := 0; k < a.GetDimension() && ok; k++ {
				delta := math.Abs(a3.GetCoord()[k] - b3.GetCoord()[k])
				ok = ok && delta < d
			}
			if ok {
				d3 := point.EuclideanDistance(*a3, *b3)
				if d3 < d {
					a, b, d = a3, b3, d3
				}
			}
		}
	}

	return a, b, d
}

func FindClosestPointPair(points []point.Point) (*point.Point, *point.Point, float64) {
	point.NumOfCalls = 0
	QuickSort(points, func(a, b point.Point) int {
		if a.GetCoord()[0] < b.GetCoord()[0] {
			return -1
		} else if a.GetCoord()[0] > b.GetCoord()[0] {
			return 1
		} else {
			return 0
		}
	})
	return fcpImpl(points)
}