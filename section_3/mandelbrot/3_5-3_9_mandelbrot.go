// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var (
	xmin, ymin, xmax, ymax = -2., -2., +2., +2.
	width, height          = 1024., 1024.
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if xMax := q.Get("x_max"); xMax != "" {
		x, err := strconv.ParseFloat(xMax, 64)
		if err != nil {
			fmt.Printf("Invalid x_max input: %v | err: %v", x, err)
		}
		xmax = x
	}

	if xMin := q.Get("x_min"); xMin != "" {
		x, err := strconv.ParseFloat(xMin, 64)
		if err != nil {
			fmt.Printf("Invalid x_min input: %v | err: %v", x, err)
		}
		xmin = x
	}

	if yMax := q.Get("y_max"); yMax != "" {
		y, err := strconv.ParseFloat(yMax, 64)
		if err != nil {
			fmt.Printf("Invalid y_max input: %v | err: %v", y, err)
		}
		width = y
	}

	if yMin := q.Get("y_min"); yMin != "" {
		y, err := strconv.ParseFloat(yMin, 64)
		if err != nil {
			fmt.Printf("Invalid y_min input: %v | err: %v", y, err)
		}
		width = y
	}
	PrintMandelbrot(w)
}

// PrintMandelbrot prints the PNG image with an io.Writer
func PrintMandelbrot(w io.Writer) {

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0.; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0.; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(int(px), int(py), mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 5 {
			return color.RGBA{25 - contrast*n, 24 - contrast*n, 245 - contrast*n, 245 - contrast*n}
		}
		if cmplx.Abs(v) > 4 {
			return color.RGBA{25 - contrast*n, 24 - contrast*n, 245 - contrast*n, 5 - contrast*n}
		}
		if cmplx.Abs(v) > 3 {
			return color.RGBA{25 - contrast*n, 234 - contrast*n, 45 - contrast*n, 5 - contrast*n}
		}
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 234 - contrast*n, 45 - contrast*n, 5 - contrast*n}
		}
	}
	return color.Black
}
