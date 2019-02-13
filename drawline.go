package bresenham

// 2016-10-22, Stéphane Bunel
//           * Do not use in production. It's just an exercise

import (
	"image/color"
	"image/draw"
)

// Floating point
func Bresenham_1(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
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
func Bresenham_2(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
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
func Bresenham_3(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
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

// Integer; remove comparison (cmp -> bit test); remove variables; float -> float * 2 * dx -> integer
func Bresenham_4(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, 2*(y2-y1)
	e, slope := dx, 2*dx
	for ; dx != 0; dx-- {
		img.Set(x1, y1, col)
		x1++
		e -= dy
		if e < 0 {
			y1++
			e += slope
		}
	}
}

// dx > dy; x1 < x2; y1 < y2
func BresenhamDxXRYD(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, 2*(y2-y1)
	e, slope := dx, 2*dx
	for ; dx != 0; dx-- {
		img.Set(x1, y1, col)
		x1++
		e -= dy
		if e < 0 {
			y1++
			e += slope
		}
	}
}

// dy > dx; x1 < x2; y1 < y2
func BresenhamDyXRYD(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := 2*(x2-x1), y2-y1
	e, slope := dy, 2*dy
	for ; dy != 0; dy-- {
		img.Set(x1, y1, col)
		y1++
		e -= dx
		if e < 0 {
			x1++
			e += slope
		}
	}
}

// dx > dy; x1 < x2; y1 > y2
func BresenhamDxXRYU(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := x2-x1, 2*(y1-y2)
	e, slope := dx, 2*dx
	for ; dx != 0; dx-- {
		img.Set(x1, y1, col)
		x1++
		e -= dy
		if e < 0 {
			y1--
			e += slope
		}
	}
}

func BresenhamDyXRYU(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	dx, dy := 2*(x2-x1), y1-y2
	e, slope := dy, 2*dy
	for ; dy != 0; dy-- {
		img.Set(x1, y1, col)
		y1--
		e -= dx
		if e < 0 {
			x1++
			e += slope
		}
	}
}

// Generalized with integer
func Bresenham(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		img.Set(x1, y1, col)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.Set(x1, y1, col)
			x1++
		}
		img.Set(x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.Set(x1, y1, col)
			y1++
		}
		img.Set(x1, y1, col)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				y1--
			}
		}
		img.Set(x1, y1, col)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.Set(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, col)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, col)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.Set(x2, y2, col)
	}
}
