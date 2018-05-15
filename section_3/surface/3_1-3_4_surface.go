// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/ajstarks/svgo"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// PrintSurface prints an SVG rendering of a 3-D surface
func PrintSurface(s *svg.SVG) {
	s.Start(width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				fmt.Println(err)
				continue
			}

			bx, by, err := corner(i, j)
			if err != nil {
				fmt.Println(err)
				continue
			}

			cx, cy, err := corner(i, j+1)
			if err != nil {
				fmt.Println(err)
				continue
			}

			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				fmt.Println(err)
				continue
			}
			s.Polygon([]int{ax, bx, cx, dx}, []int{ay, by, cy, dy}, "stroke: grey; fill: white; stroke-width: 0.7")
		}
	}
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query()
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	PrintSurface(s)
	s.End()
}

func corner(i, j int) (int, int, error) {
	var err error
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) {
		err = fmt.Errorf("invalid polygon: infinite float x: %g, y: %g", x, y)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return int(sx), int(sy), err
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
