package main

import (
	"fmt"
	"log"
	"runtime"
	"tucil/stima/pairit/algorithm"
	"tucil/stima/pairit/point"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
	SCREEN_DEPTH  = 500
)

type (
	getGlParam func(uint32, uint32, *int32)
	getInfoLog func(uint32, int32, *int32, *uint8)
)

func checkGlError(glObject uint32, errorParam uint32, getParamFn getGlParam,
	getInfoLogFn getInfoLog, failMsg string,
) {
	var success int32
	getParamFn(glObject, errorParam, &success)
	if success != 1 {
		var infoLog [512]byte
		getInfoLogFn(glObject, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))
		log.Fatalln(failMsg, "\n", string(infoLog[:512]))
	}
}

func checkShaderCompileErrors(shader uint32) {
	checkGlError(shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"ERROR::SHADER::COMPILE_FAILURE")
}

func checkProgramLinkErrors(program uint32) {
	checkGlError(program, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog,
		"ERROR::PROGRAM::LINKING_FAILURE")
}

var vertexShaderSource = `
#version 410
layout (location = 0) in vec3 position;
void main() {
	gl_Position = vec4(position.x, position.y, position.z, 1.0);
}
`

var fragmentShaderSource = `
#version 410
uniform vec3 color;
void main()
{
	gl_FragColor = vec4(color, 1.0f);
}
`

func compileShaders() []uint32 {
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	shaderSourceChars, freeVertexShaderFunc := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, shaderSourceChars, nil)
	gl.CompileShader(vertexShader)
	checkShaderCompileErrors(vertexShader)

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	shaderSourceChars, freeFragmentShaderFunc := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, shaderSourceChars, nil)
	gl.CompileShader(fragmentShader)
	checkShaderCompileErrors(fragmentShader)

	defer freeFragmentShaderFunc()
	defer freeVertexShaderFunc()

	return []uint32{vertexShader, fragmentShader}
}

func linkShaders(shaders []uint32) uint32 {
	program := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
	}
	gl.LinkProgram(program)
	checkProgramLinkErrors(program)

	// shader objects are not needed after they are linked into a program object
	for _, shader := range shaders {
		gl.DeleteShader(shader)
	}

	return program
}

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "PairIt", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	shaders := compileShaders()
	shaderProgram := linkShaders(shaders)

	gl.PointSize(10)

	points := []point.Point3D{
		point.CreatePoint3D(-200, 30, 53),
		point.CreatePoint3D(169, 44, 124),
		point.CreatePoint3D(130, -143, -34),
		point.CreatePoint3D(-261, 33, 211),
		point.CreatePoint3D(236, -58, 29),
		point.CreatePoint3D(71, 19, 99),
		point.CreatePoint3D(-53, 79, 27),
		point.CreatePoint3D(-38, -125, 11),
		point.CreatePoint3D(99, -85, 40),
		point.CreatePoint3D(120, 174, -34),
	}

	p1, p2, d := algorithm.FindClosestPoint3DPair(points)

	fmt.Println("===== List of Points =====")
	for i := 0; i < len(points); i++ {
		fmt.Printf("%v at address %p\n", points[i], &points[i])
	}
	fmt.Println("===== Closest Pair =====")
	fmt.Printf("%v at address %p\n", *p1, p1)
	fmt.Printf("%v at address %p\n", *p2, p2)
	fmt.Println("Distance:", d)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(shaderProgram)

		colorUniformLocation := gl.GetUniformLocation(shaderProgram, gl.Str("color\x00"))

		for i := 0; i < len(points); i++ {
			if p1 == &points[i] || p2 == &points[i] {
				gl.Uniform3f(colorUniformLocation, 0.8, 0, 0)
			} else {
				gl.Uniform3f(colorUniformLocation, 0, 0.8, 0.8)
			}
			vao := createPointVAO(points[i])
			gl.BindVertexArray(vao)
			gl.DrawArrays(gl.POINTS, 0, 1)
			gl.BindVertexArray(0)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func normalizedPoint(p point.Point3D) []float32 {
	x := p.GetX() / SCREEN_WIDTH * 2
	y := p.GetY() / SCREEN_HEIGHT * 2
	z := p.GetZ() / SCREEN_DEPTH * 2
	return []float32{x, y, z}
}

func createPointVAO(p point.Point3D) uint32 {
	vertices := normalizedPoint(p)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 0, 0)

	gl.BindVertexArray(0)

	return vao
}
