package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"tucil/stima/pairit/algorithm"
	. "tucil/stima/pairit/point"

	. "github.com/klauspost/cpuid/v2"
)

type fcpFunction func([]Point) (*Point, *Point, float64)

var (
	dim, n     int
	upperBound float64
	points     []Point
	p1, p2     *Point
	plotData   *os.File
)

func main() {
	// handle SIGTERM and SIGINT for delete temporary file (plotData)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		deleteTempFile()
		os.Exit(1)
	}()
	defer deleteTempFile()

	printBanner()

inputDim:
	fmt.Print("Dimension\t: ")
	_, err := fmt.Scanf("%d", &dim)
	if err != nil || dim < 2 {
		fmt.Println("Invalid input!")
		goto inputDim // Is it bad practice? I don't think so
	}

inputN:
	fmt.Print("Number of points: ")
	_, err = fmt.Scanf("%d", &n)
	if err != nil || n < 2 {
		fmt.Println("Invalid input!")
		goto inputN
	}

	points = generatePoints(dim, n, float64(n))
	performFcpAlgorithm("Divide and Conquer", algorithm.FindClosestPointPair)
	performFcpAlgorithm("Brute Force", algorithm.BruteForceFCP)

	path, err := exec.LookPath("gnuplot")
	if dim <= 3 {
		if err != nil {
			fmt.Println("Gnuplot not found in your PATH. Visualization is not performed")
		} else {
			if dim == 3 {
				process3D(points)
				runGnuplot(path)
			} else {
				process2D(points)
				runGnuplot(path)
			}
		}
	}
}

func performFcpAlgorithm(title string, algo fcpFunction) {
	NumOfCalls = 0
	var d float64
	start := time.Now()
	p1, p2, d = algo(points)
	executionTime := time.Since(start)
	fmt.Println()
	fmt.Printf("\x1b[93m====== %s ======\x1b[0m\n", title)
	fmt.Printf("Point 1\t\t: %v\n", p1.GetCoord())
	fmt.Printf("Point 2\t\t: %v\n", p2.GetCoord())
	fmt.Printf("Distance\t: %f\n", d)
	fmt.Printf("Execution time\t: %.9f s (%s)\n", float64(executionTime.Nanoseconds())/1e9, CPU.BrandName)
	fmt.Printf("The Euclidean distance function is called %dx\n", NumOfCalls)
}

func runGnuplot(path string) {
	var proc *exec.Cmd
	var format string
	if dim == 3 {
		format = "splot \"%s\" u 1:2:3:4 t \"\" w p pt 7 ps 1 lc variable"
	} else {
		format = "plot \"%s\" u 1:2:3 t \"\" w p pt 7 ps 1 lc variable"
	}
	commands := []string{
		"set term qt title 'PairIt'",
		fmt.Sprintf("set xrange [0:%f]", upperBound),
		fmt.Sprintf("set yrange [0:%f]", upperBound),
	}
	if dim == 3 {
		commands = append(commands, fmt.Sprintf("set zrange [0:%f]", upperBound))
	}
	commands = append(commands,
		fmt.Sprintf(format, plotData.Name()),
		"pause mouse close",
		"q",
	)
	if runtime.GOOS != "windows" {
		proc = exec.Command(path)
		stdin, _ := proc.StdinPipe()
		if err := proc.Start(); err != nil {
			err_str := fmt.Sprintln("**", err.Error())
			panic(err_str)
		}
		for _, s := range commands {
			plotCmd(stdin, s)
		}
	} else {
		proc = exec.Command(path, append(commands, "-e")...)
		if err := proc.Start(); err != nil {
			err_str := fmt.Sprintln("**", err.Error())
			panic(err_str)
		}
	}
	if err := proc.Wait(); err != nil {
		err_str := fmt.Sprintln("**", err.Error())
		panic(err_str)
	}
}

func process3D(points []Point) {
	var err error
	plotData, err = os.CreateTemp(".", "gnuplot-data-")
	if err != nil {
		err_str := fmt.Sprintln("**", err.Error())
		panic(err_str)
	}
	for i := 0; i < len(points); i++ {
		x := points[i].GetCoord()[0]
		y := points[i].GetCoord()[1]
		z := points[i].GetCoord()[2]
		color := 0
		if &points[i] == p1 || &points[i] == p2 {
			color = 7
		}
		line := fmt.Sprintf("%v %v %v %v\n", x, y, z, color)
		if _, err := plotData.WriteString(line); err != nil {
			err_str := fmt.Sprintln("**", err.Error())
			panic(err_str)
		}
	}
	plotData.Close()
}

func process2D(points []Point) {
	var err error
	plotData, err = os.CreateTemp(".", "gnuplot-data-")
	if err != nil {
		err_str := fmt.Sprintln("**", err.Error())
		panic(err_str)
	}
	for i := 0; i < len(points); i++ {
		x := points[i].GetCoord()[0]
		y := points[i].GetCoord()[1]
		color := 0
		if &points[i] == p1 || &points[i] == p2 {
			color = 7
		}
		line := fmt.Sprintf("%v %v %v\n", x, y, color)
		if _, err := plotData.WriteString(line); err != nil {
			err_str := fmt.Sprintln("**", err.Error())
			panic(err_str)
		}
	}
	plotData.Close()
}

func plotCmd(stdin io.Writer, command string) {
	if _, err := io.WriteString(stdin, command+"\n"); err != nil {
		fmt.Println("**", err.Error())
	}
}

func generatePoints(dim, n int, ub float64) []Point {
	rand.Seed(time.Now().UnixNano())
	upperBound = ub
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		coord := make([]float64, dim)
		for j := 0; j < dim; j++ {
			coord[j] = rand.Float64() * upperBound
		}
		points[i] = CreatePoint(coord...)
	}
	return points
}

func printBanner() {
	if banner, err := os.ReadFile("banner"); err != nil {
		fmt.Println("**", err.Error())
	} else {
		fmt.Println(string(banner))
	}
}

func deleteTempFile() {
	if plotData != nil {
		os.Remove(plotData.Name())
	}
}
