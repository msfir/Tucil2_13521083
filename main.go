package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"tucil/stima/pairit/algorithm"
	"tucil/stima/pairit/point"

	. "github.com/klauspost/cpuid/v2"
)

var upperBound float64

func generatePoints(n int, ub float64) []point.Point3D {
	upperBound = ub
	points := make([]point.Point3D, n)
	for i := 0; i < n; i++ {
		x := rand.Float64() * upperBound
		y := rand.Float64() * upperBound
		z := rand.Float64() * upperBound
		points[i] = point.CreatePoint3D(x, y, z)
	}
	return points
}

func main() {
	var n int

input:
	fmt.Print("Number of points: ")
	_, err := fmt.Scan(&n)
	if err != nil || n < 2 {
		fmt.Println("Invalid input!")
		goto input // Is it bad practice? I don't think so
	}

	points := generatePoints(n, float64(n))

	start := time.Now()
	p1, p2, d := algorithm.FindClosestPoint3DPair(points)
	executionTime := time.Since(start)
	fmt.Println("====== Closest Pair ======")
	fmt.Printf("Point 1: (%v, %v %v)\n", p1.GetX(), p1.GetY(), p1.GetZ())
	fmt.Printf("Point 2: (%v, %v %v)\n", p2.GetX(), p2.GetY(), p2.GetZ())
	fmt.Printf("Distance: %f\n", d)
	fmt.Printf("Execution time: %f s (%s)\n", float64(executionTime.Nanoseconds())/1e9, CPU.BrandName)
	fmt.Printf("The Euclidean distance function is called %d times\n", point.NumOfCalls)

	path, err := exec.LookPath("gnuplot")
	if err != nil {
		fmt.Printf("** %v\n", err)
	} else {
		data, err := os.CreateTemp(".", "gnuplot-data-")
		defer os.Remove(data.Name())
		if err != nil {
			err_str := fmt.Sprintf("** %v\n", err)
			panic(err_str)
		}
		for i := 0; i < n; i++ {
			x := points[i].GetX()
			y := points[i].GetY()
			z := points[i].GetZ()
			color := 0
			if &points[i] == p1 || &points[i] == p2 {
				color = 7
			}
			line := fmt.Sprintf("%v %v %v %v\n", x, y, z, color)
			if _, err := data.WriteString(line); err != nil {
				err_str := fmt.Sprintf("** %v\n", err)
				panic(err_str)
			}
		}
		data.Close()
		proc := exec.Command(path)
		stdin, _ := proc.StdinPipe()
		if proc.Start() != nil {
			err_str := fmt.Sprintf("** %v\n", err)
			panic(err_str)
		}
		cmd := fmt.Sprintf("splot \"%s\" u 1:2:3:4 t \"\" w p pt 7 ps 1 lc variable", data.Name())
		plotCmd(stdin, "set term qt title 'PairIt'")
		plotCmd(stdin, fmt.Sprintf("set xrange [0:%f]", upperBound))
		plotCmd(stdin, fmt.Sprintf("set yrange [0:%f]", upperBound))
		plotCmd(stdin, fmt.Sprintf("set zrange [0:%f]", upperBound))
		plotCmd(stdin, cmd)
		plotCmd(stdin, "pause mouse close")
		plotCmd(stdin, "q")
		if proc.Wait() != nil {
			err_str := fmt.Sprintf("** %v\n", err)
			panic(err_str)
		}
	}
}

func plotCmd(stdin io.Writer, command string) {
	if _, err := io.WriteString(stdin, command+"\n"); err != nil {
		fmt.Printf("** %v\n", err)
	}
}
