package ch3

// surface Calcs a render SVG 3D surface `go run book.go > out.svg`
import (
	"fmt"
	"math"
	"os"
)

type tFunc func(x, y float64) float64

// ShowSurfaceEx3_2 is | `>_ go run book.go eggbox > outEx3.2.svg` OR `>_ go run book.go saddle > outEx3.2.svg`
func ShowSurfaceEx3_2() {
	usage := "Usage: Ex3.2 saddle|eggbox"
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	var f tFunc
	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	default:
		fmt.Println(usage)
		os.Exit(1)
	}

	svg(f)
}

func svg(f tFunc) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill: white; strokewidth:0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := cornerTFunc(i+1, j, f)
			bx, by := cornerTFunc(i, j, f)
			cx, cy := cornerTFunc(i, j+1, f)
			dx, dy := cornerTFunc(i+1, j+1, f)

			// Exercise 3.1 solution
			corners := []float64{ax, ay, bx, by, cx, cy, dx, dy}
			if any(corners, func(c float64) bool {
				return math.IsNaN(c)
			}) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g, %g,%g, %g,%g, %g,%g' /> \n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func cornerTFunc(i, j int, f tFunc) (float64, float64) {
	// Encontra o ponto (x,y) no canto da celula (i,j)
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)

	// calcula a superficie z da superficie
	z := f(x, y)

	// Make a isometric projection isometrica de (x,y,z) over (sx,sy) canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func any(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
