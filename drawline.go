package bresenham

// 2016-10-22, Stéphane Bunel
//           * Do not use in production. Not tested.

import (
	"image"
	"image/color"
)

// Floating point
func Bresenham_1(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, y2-y1
	a := float64(dy) / float64(dx)
	b := int(float64(y1) - a*float64(x1))

	img.Set(x1, y1, col)
	for x := x1 + 1; x <= x2; x++ {
		y := int(a*float64(x)) + b
		img.Set(x, y, col)
	}
}

// Floating point with error accumulator
func Bresenham_2(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, y2-y1
	a := float64(dy) / float64(dx)
	e, e_max, e_sub := 0.0, 0.5, 1.0
	y := y1

	img.Set(x1, y1, col)
	for x := x1 + 1; x <= x2; x++ {
		img.Set(x, y, col)
		e += a
		if e > e_max {
			y += 1
			e -= e_sub
		}
	}
}

// Integer float -> float * dx -> integer
func Bresenham_3(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, y2-y1
	// e, e_max, e_sub := 0*dx, dx/2, dx
	e, e_max, e_sub := dx, dx>>1, dx
	y := y1

	img.Set(x1, y1, col)
	for x := x1 + 1; x <= x2; x++ {
		img.Set(x, y, col)
		e += dy // <= dy/dx * dx
		if e > e_max {
			y += 1
			e -= e_sub
		}
	}
}

// Integer float -> float * 2 * dx -> integer
func Bresenham_4(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, 2*(y2-y1)
	// e, e_max, e_sub := 2*0*dx, dx, 2*dx
	e, e_max, e_sub := 0, dx, dx<<1
	y := y1

	img.Set(x1, y1, col)
	for x := x1 + 1; x <= x2; x++ {
		img.Set(x, y, col)
		e += dy // <= 2 * (dy/dx * dx)
		if e > e_max {
			y += 1
			e -= e_sub
		}
	}
}

// Generalized with integer
func Bresenham(img *image.RGBA, p1, p2 image.Point, col color.Color) {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y

	// Alway plot first point
	img.Set(x1, y1, col)

	// Line is a point ?
	if x1 == x2 && y1 == y2 {
		return
	}

	// draw a vertical line ?
	if x1 == x2 {
		if y1 < y2 {
			y1++
			for y1 <= y2 {
				img.Set(x1, y1, col)
				y1++
			}
		} else {
			y2++
			for y2 <= y1 {
				img.Set(x1, y2, col)
				y2++
			}
		}
		return
	}

	// draw a horizontal line ?
	if y1 == y2 {
		if x1 < x2 {
			x1++
			for x1 <= x2 {
				img.Set(x1, y1, col)
				x1++
			}
		} else {
			x2++
			for x2 <= x1 {
				img.Set(x2, y1, col)
				x2++
			}
		}
		return
	}

	var dx, dy, xsign, ysign int

	if x1 < x2 {
		dx = x2 - x1
		xsign = +1
	} else {
		dx = x1 - x2
		xsign = -1
	}

	if y1 < y2 {
		dy = y2 - y1
		ysign = +1
	} else {
		dy = y1 - y2
		ysign = -1
	}

	// a == 1 ?
	if dx == dy {
		for x1 != x2 {
			x1 += xsign
			y1 += ysign
			img.Set(x1, y1, col)
		}
		return
	}

	// INFO: e, e_max, e_sub := 2*0*dx, dx, 2*dx

	if dx > dy {
		e, e_max, e_sub, y := 0, dx, dx<<1, y1
		for x1 != x2 {
			x1 += xsign
			img.Set(x1, y, col)
			e += dy << 1 // <= 2 * ( dy/dx * dx )
			if e > e_max {
				y += ysign
				e -= e_sub
			}
		}
	} else {
		e, e_max, e_sub, x := 0, dy, dy<<1, x1
		for y1 != y2 {
			y1 += ysign
			img.Set(x, y1, col)
			e += dx << 1 // <= 2 * ( dy/dx * dx )
			if e > e_max {
				x += xsign
				e -= e_sub
			}
		}
	}
	img.Set(x2, y2, col)
}
