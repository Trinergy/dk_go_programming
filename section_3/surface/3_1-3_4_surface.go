// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/ajstarks/svgo"
)

var (
	width, height = 600.0, 320.0        // canvas size in pixels
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
	style   = "stroke: grey; fill: white; stroke-width: 0.7"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// PrintSurface prints an SVG rendering of a 3-D surface
func PrintSurface(s *svg.SVG) {
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
			xs := []int{ax, bx, cx, dx}
			ys := []int{ay, by, cy, dy}
			customStyle := "stroke: blue; fill: red; stroke-width: 0.7"
			useCustomStyle := false

			for i := range ys {
				if ys[i] < 80 {
					useCustomStyle = true
					break
				}
			}

			if useCustomStyle {
				s.Polygon(xs, ys, customStyle)
			} else {
				s.Polygon(xs, ys, style)
			}
		}
	}
}

// handler prints the SVG of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	q := r.URL.Query()
	if queryWidth := q.Get("width"); queryWidth != "" {
		w, err := strconv.ParseFloat(queryWidth, 64)
		if err != nil {
			fmt.Printf("Invalid Width input: %v | err: %v", w, err)
		}
		width = w
	}

	if queryHeight := q.Get("height"); queryHeight != "" {
		h, err := strconv.ParseFloat(queryHeight, 64)
		if err != nil {
			fmt.Printf("Invalid Height input: %v | err: %v", h, err)
		}
		height = h
	}

	s := svg.New(w)
	s.Start(int(width), int(height))
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
