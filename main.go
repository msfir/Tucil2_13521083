package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"tucil/stima/pairit/point"
)

const SIZE = 100

func generatePoints(n int) []point.Point3D {
	points := make([]point.Point3D, n)
	for i := 0; i < n; i++ {
		x := (rand.Float32() - 0.5) * SIZE
		y := (rand.Float32() - 0.5) * SIZE
		z := (rand.Float32() - 0.5) * SIZE
		points[i] = point.CreatePoint3D(x, y, z)
	}
	return points
}

func main() {
	n := 100

	points := generatePoints(n)

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
		cmd := fmt.Sprintf("splot \"%s\" u 1:2:3:4 t \"points\" w p pt 7 ps 1 lc variable", data.Name())
		plotCmd(stdin, fmt.Sprintf("set xrange [-%[1]d:%[1]d]", SIZE/2))
		plotCmd(stdin, fmt.Sprintf("set yrange [-%[1]d:%[1]d]", SIZE/2))
		plotCmd(stdin, fmt.Sprintf("set zrange [-%[1]d:%[1]d]", SIZE/2))
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
