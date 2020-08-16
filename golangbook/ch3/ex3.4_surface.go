package ch3

// surface Calcs a render SVG 3D surface `go run book.go > out.svg`
import (
	"fmt"
	"math"
	"os"
)

// ShowSurfaceEx3_4 is | `>_ go run book.go eggbox > outEx3.2.svg` OR `>_ go run book.go saddle > outEx3.2.svg`
func ShowSurfaceEx3_4() {
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

	svgColor(f)
}

func svgColor(f tFunc) {
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
