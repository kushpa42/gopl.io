// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
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
	//!-main
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			surface(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	surface(os.Stdout)
}

func surface(out io.Writer) {
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	out.Write([]byte(s))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)

			if err != nil {
				continue
			}

			bx, by, err := corner(i, j)

			if err != nil {
				continue
			}

			cx, cy, err := corner(i, j+1)

			if err != nil {
				continue
			}

			dx, dy, err := corner(i+1, j+1)

			if err != nil {
				continue
			}

			color := color(i, j)

			s := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s; stroke-width: 0.7'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)

			out.Write([]byte(s))
		}
	}

	out.Write([]byte("</svg>"))
}

func color(i, j int) string {
	_, _, z := coordinates(i, j)

	if z > 0 {
		return "#ff0000"
	}

	return "#0000ff"
}

func coordinates(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	return x, y, z
}

func corner(i, j int) (float64, float64, error) {
	x, y, z := coordinates(i, j)

	if math.IsInf(z, 0) {
		return 0, 0, errors.New("infinite number")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	return math.Sin(r) / r
}

//!-
