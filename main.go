package main

import (
	"fmt"
	"tucil/stima/pairit/point"
)

func main() {
	p1 := point.CreatePoint3D(1, 2, 3)
	p2 := point.CreatePoint3D(5, 7, 2)
	fmt.Println(p1.ToSlice())
	fmt.Println("Hello World", point.Distance(p1, p2))
}
