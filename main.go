package main

import (
	"fmt"
	. "tucil/stima/pairit/algorithm"
	"tucil/stima/pairit/point"
)

func main() {
	p1 := point.CreatePoint3D(1, 5, 8)
	p2 := point.CreatePoint3D(3, 7, 2)
	p3 := point.CreatePoint3D(-4, 5, 7)
	p4 := point.CreatePoint3D(4, 9, 11)
	p5 := point.CreatePoint3D(5, 12, -4)
	p6 := point.CreatePoint3D(6, 1, -8)
	p7 := point.CreatePoint3D(3, 3, 4)
	points := []point.Point3D{p1, p2, p3, p4, p5, p6, p7}
	a, b, d := FindClosestPoint3DPair(points)
	for i := 0; i < len(points); i++ {
		fmt.Printf("%v at adress %p\n", points[i], &points[i])
	}
	fmt.Println("===== Closest Pair =====")
	fmt.Printf("Point 1: %v at address %p\n", *a, a)
	fmt.Printf("Point 2: %v at address %p\n", *b, b)
	fmt.Printf("Distance: %f\n", d)
}
