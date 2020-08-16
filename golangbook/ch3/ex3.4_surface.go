package ch3

// surface Calcs a render SVG 3D surface `go run book.go > out.svg`
import (
	"fmt"
	"io"
	"math"
	"net/http"
)

// ShowSurfaceEx3_4 is
func ShowSurfaceEx3_4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svgWriter(w)
	})

}

func svgWriter(w io.Writer) {
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
