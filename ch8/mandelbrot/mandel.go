// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

type Job struct {
	px int
	py int
	z  complex128
}

type Result struct {
	px int
	py int
	c  color.Color
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job.px, job.py, mandelbrot(job.z)}
		results <- output
	}
	wg.Done()
}

func createWorkPool(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate() {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			jobs <- Job{px, py, z}
		}
	}
	close(jobs)
}
func createImg(nWorkers int) {
	go allocate()
	go createWorkPool(nWorkers)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for result := range results {
		img.Set(result.px, result.py, result.c)
	}

	//png.Encode(os.Stdout, img)
}

var nWorkers = flag.Int("n", 10, "numbers of worker")

func main() {
	flag.Parse()
	start := time.Now()
	createImg(*nWorkers)
	fmt.Fprintf(os.Stdout, "%d workers time: %v\n", *nWorkers, time.Since(start))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	time.Sleep(1 * time.Nanosecond)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
