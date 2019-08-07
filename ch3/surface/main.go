// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

import (
	"github.com/ajstarks/svgo"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

const defaultstyle = "fill:rgb(127,0,0)"

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	if (len(os.Args) > 1) && (os.Args[1] == "web") {

		http.HandleFunc("/surface", surface)
		http.HandleFunc("/circle", circle)

		log.Fatal(http.ListenAndServe(":2003", nil))
	}
	surfacesvg(os.Stdout)

}

func shapestyle(path string) string {
	i := strings.LastIndex(path, "/") + 1
	if i > 0 && len(path[i:]) > 0 {
		return "fill:" + path[i:]
	}
	return defaultstyle
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	circlesvg(w)
}

func circlesvg(out io.Writer) {
	s := svg.New(out)
	s.Start(500, 500)
	s.Title("Circle")
	s.Circle(250, 250, 125, defaultstyle)
	s.End()
}
func surface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surfacesvg(w)
}

func surfacesvg(out io.Writer) {
	keypoints := GradientTable{
		{MustParseHex("#ff0000"), -0.3},
		{MustParseHex("#0000ff"), 0.5},
	}

	fmt.Fprintln(out, "<?xml version='1.0'?>")
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			h := (az + bz + cz + dz) / 4
			c := keypoints.GetInterpolatedColorFor(h)

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.Hex())
		}
	}
	fmt.Fprintln(out, "</svg>")
}
func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, err := f(x, y)
	if !err {
		z = 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sz := z

	return sx, sy, sz
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	if math.IsNaN(r) {
		return 0, false
	}
	return math.Sin(r) / r, true
}

//!-
