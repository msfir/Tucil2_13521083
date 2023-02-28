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

func generatePoints(n int, ub float64) []point.Point {
	upperBound = ub
	points := make([]point.Point, n)
	for i := 0; i < n; i++ {
		x := rand.Float64() * upperBound
		y := rand.Float64() * upperBound
		z := rand.Float64() * upperBound
		points[i] = point.CreatePoint(x, y, z)
	}
	return points
}

func main() {
	printBanner()

	var n int

input:
	fmt.Print("Number of points: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil || n < 2 {
		fmt.Println("Invalid input!")
		goto input // Is it bad practice? I don't think so
	}

	points := generatePoints(n, float64(n))

	start := time.Now()
	p1, p2, d := algorithm.FindClosestPointPair(points)
	executionTime := time.Since(start)
	fmt.Println("\x1b[93m====== Divide and Conquer ======\x1b[0m")
	fmt.Printf("Point 1\t\t: (%v, %v, %v)\n", p1.GetCoord()[0], p1.GetCoord()[1], p1.GetCoord()[2])
	fmt.Printf("Point 2\t\t: (%v, %v, %v)\n", p2.GetCoord()[0], p2.GetCoord()[1], p2.GetCoord()[2])
	fmt.Printf("Distance\t: %f\n", d)
	fmt.Printf("Execution time\t: %.9f s (%s)\n", float64(executionTime.Nanoseconds())/1e9, CPU.BrandName)
	fmt.Printf("The Euclidean distance function is called %dx\n", point.NumOfCalls)

	point.NumOfCalls = 0
	fmt.Println()

	start = time.Now()
	p1, p2, d = algorithm.BruteForceFCP(points)
	executionTime = time.Since(start)
	fmt.Println("\x1b[93m====== Brute Force ======\x1b[0m")
	fmt.Printf("Point 1\t\t: (%v, %v, %v)\n", p1.GetCoord()[0], p1.GetCoord()[1], p1.GetCoord()[2])
	fmt.Printf("Point 2\t\t: (%v, %v, %v)\n", p2.GetCoord()[0], p2.GetCoord()[1], p2.GetCoord()[2])
	fmt.Printf("Distance\t: %f\n", d)
	fmt.Printf("Execution time\t: %.9f s (%s)\n", float64(executionTime.Nanoseconds())/1e9, CPU.BrandName)
	fmt.Printf("The Euclidean distance function is called %dx\n", point.NumOfCalls)

	path, err := exec.LookPath("gnuplot")
	if err != nil {
		fmt.Printf("** %v\n", err.Error())
	} else {
		data, err := os.CreateTemp(".", "gnuplot-data-")
		defer os.Remove(data.Name())
		if err != nil {
			err_str := fmt.Sprintf("** %v\n", err.Error())
			panic(err_str)
		}
		for i := 0; i < n; i++ {
			x := points[i].GetCoord()[0]
			y := points[i].GetCoord()[1]
			z := points[i].GetCoord()[2]
			color := 0
			if &points[i] == p1 || &points[i] == p2 {
				color = 7
			}
			line := fmt.Sprintf("%v %v %v %v\n", x, y, z, color)
			if _, err := data.WriteString(line); err != nil {
				err_str := fmt.Sprintf("** %v\n", err.Error())
				panic(err_str)
			}
		}
		data.Close()
		proc := exec.Command(path)
		stdin, _ := proc.StdinPipe()
		if proc.Start() != nil {
			err_str := fmt.Sprintf("** %v\n", err.Error())
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
			err_str := fmt.Sprintf("** %s\n", err.Error())
			panic(err_str)
		}
	}
}

func plotCmd(stdin io.Writer, command string) {
	if _, err := io.WriteString(stdin, command+"\n"); err != nil {
		fmt.Printf("** %s\n", err.Error())
	}
}

func printBanner() {
	if banner, err := os.ReadFile("banner"); err != nil {
		fmt.Println("** ", err.Error())
	} else {
		fmt.Println(string(banner))
	}
}
