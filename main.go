package main

import (
	"fmt"
	. "tucil/stima/pairit/algorithm"
	"tucil/stima/pairit/point"
)

func compareX(a point.Point3D, b point.Point3D) int {
	if a.GetX() < b.GetX() {
		return -1
	} else if a.GetX() > b.GetX() {
		return 1
	} else {
		return 0
	}
}

func main() {
	p1 := point.CreatePoint3D(5, 2, 3)
	p2 := point.CreatePoint3D(4, 7, 2)
	p3 := point.CreatePoint3D(8, 3, 3)
	p4 := point.CreatePoint3D(3, 3, 3)
	p5 := point.CreatePoint3D(2, 3, 3)
	p6 := point.CreatePoint3D(0, 3, 3)
	points := []point.Point3D{p1, p2, p3, p4, p5, p6}
	fmt.Println("Before sorting: ", points)
	QuickSort(points, compareX)
	fmt.Println("After sorting: ", points)
}
