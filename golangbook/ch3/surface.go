package ch3

// surface Calcs a render SVG 3D surface `go run book.go > out.svg`
import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // Canvas's size in pixel
	cells         = 100                 // cells numbers grade
	xyrange       = 30.0                // axes intervals (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // unit per pixels x or y
	zscale        = height * .4         //unit pixels z
	angle         = math.Pi / 6         // axes's angle x,y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30°), cosseno(30°)

// ShowSurface is
func ShowSurface() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill: white; strokewidth:0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
	OUTER:
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// Exercise 3.1 solution
			corners := []float64{ax, ay, bx, by, cx, cy, dx, dy}
			for _, cElement := range corners {
				if math.IsNaN(cElement) {
					continue OUTER
				}
			}

			fmt.Printf("<polygon points='%g,%g, %g,%g, %g,%g, %g,%g' /> \n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
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

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance of (0,0)
	return math.Sin(r) / r
}
