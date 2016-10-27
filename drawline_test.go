package bresenham

// 2016-10-22, Stéphane Bunel
//           * go test
//           * go test -bench=.

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

var colRED = color.RGBA{255, 0, 0, 255}
var colGREEN = color.RGBA{0, 255, 0, 255}
var colBLUE = color.RGBA{0, 0, 255, 255}
var colWHITE = color.RGBA{255, 255, 255, 255}

func drawCross(img *image.RGBA, x, y int, col color.Color) {
	for s := -3; s < 4; s++ {
		img.Set(x+s, y, col)
		img.Set(x, y+s, col)
	}
}

func Test_BresenhamDxXRYD(t *testing.T) {
	var imgRect = image.Rect(0, 0, 100, 100)
	var img = image.NewRGBA(imgRect)

	drawCross(img, imgRect.Min.X, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Max.Y-1, colWHITE)
	drawCross(img, imgRect.Min.X, imgRect.Max.Y-1, colWHITE)

	x1, y1 := 17, 11
	drawCross(img, x1, y1, colRED)
	x2, y2 := 71, 41
	drawCross(img, x2, y2, colGREEN)

	BresenhamDxXRYD(img, x1, y1, x2, y2, colBLUE)

	// Save result
	filename := "bresenhamDxXRYD.png"
	toimg, _ := os.Create(filename)
	defer toimg.Close()
	png.Encode(toimg, img)
}

func Test_BresenhamDyXRYD(t *testing.T) {
	var imgRect = image.Rect(0, 0, 100, 100)
	var img = image.NewRGBA(imgRect)

	drawCross(img, imgRect.Min.X, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Max.Y-1, colWHITE)
	drawCross(img, imgRect.Min.X, imgRect.Max.Y-1, colWHITE)

	x1, y1 := 17, 11
	drawCross(img, x1, y1, colRED)
	x2, y2 := 47, 71
	drawCross(img, x2, y2, colGREEN)

	BresenhamDyXRYD(img, x1, y1, x2, y2, colBLUE)

	// Save result
	filename := "bresenhamDyXRYD.png"
	toimg, _ := os.Create(filename)
	defer toimg.Close()
	png.Encode(toimg, img)
}

func Test_BresenhamDxXRYU(t *testing.T) {
	var imgRect = image.Rect(0, 0, 100, 100)
	var img = image.NewRGBA(imgRect)

	drawCross(img, imgRect.Min.X, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Max.Y-1, colWHITE)
	drawCross(img, imgRect.Min.X, imgRect.Max.Y-1, colWHITE)

	x1, y1 := 11, 45
	drawCross(img, x1, y1, colRED)

	x2, y2 := 71, 7
	drawCross(img, x2, y2, colGREEN)

	BresenhamDxXRYU(img, x1, y1, x2, y2, colBLUE)

	// Save result
	filename := "bresenhamDxXRYU.png"
	toimg, _ := os.Create(filename)
	defer toimg.Close()
	png.Encode(toimg, img)
}

func Test_BresenhamDyXRYU(t *testing.T) {
	var imgRect = image.Rect(0, 0, 100, 100)
	var img = image.NewRGBA(imgRect)

	drawCross(img, imgRect.Min.X, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Max.Y-1, colWHITE)
	drawCross(img, imgRect.Min.X, imgRect.Max.Y-1, colWHITE)

	x1, y1 := 24, 71
	drawCross(img, x1, y1, colRED)
	x2, y2 := 47, 11
	drawCross(img, x2, y2, colGREEN)

	BresenhamDyXRYU(img, x1, y1, x2, y2, colBLUE)

	// Save result
	filename := "bresenhamDyXRYU.png"
	toimg, _ := os.Create(filename)
	defer toimg.Close()
	png.Encode(toimg, img)
}

func Test_Bresenham(t *testing.T) {
	var imgRect = image.Rect(0, 0, 100, 100)
	var img = image.NewRGBA(imgRect)
	var x1, y1, x2, y2 int

	drawCross(img, imgRect.Min.X, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Min.Y, colWHITE)
	drawCross(img, imgRect.Max.X-1, imgRect.Max.Y-1, colWHITE)
	drawCross(img, imgRect.Min.X, imgRect.Max.Y-1, colWHITE)

	// H line
	x1, y1, x2, y2 = 50, 20, 90, 20
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// V line
	x1, y1, x2, y2 = 70, 10, 70, 40
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// m=1 line
	x1, y1, x2, y2 = 60, 60, 90, 90
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// dxXLYD line
	x1, y1, x2, y2 = 45, 10, 4, 20
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// dxXRYU line
	x1, y1, x2, y2 = 10, 80, 60, 70
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// dyXRYD line
	x1, y1, x2, y2 = 12, 10, 44, 90
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// dyXLYU line
	x1, y1, x2, y2 = 30, 95, 8, 30
	drawCross(img, x1, y1, colRED)
	drawCross(img, x2, y2, colGREEN)
	Bresenham(img, x1, y1, x2, y2, colBLUE)

	// Save result
	filename := "bresenhamALL.png"
	toimg, _ := os.Create(filename)
	defer toimg.Close()
	png.Encode(toimg, img)
}

// ----- bench part -----

var imgRect = image.Rect(0, 0, 100, 100)
var img = image.NewRGBA(imgRect)
var x1, y1, x2, y2 = 10, 10, 90, 25

func Benchmark_Bresenham_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_1(img, x1, y1, x2, y2, colWHITE)
	}
}

func Benchmark_Bresenham_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_2(img, x1, y1, x2, y2, colWHITE)
	}
}

func Benchmark_Bresenham_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_3(img, x1, y1, x2, y2, colWHITE)
	}
}

func Benchmark_Bresenham_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_4(img, x1, y1, x2, y2, colWHITE)
	}
}

func Benchmark_BresenhamDxXRYD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BresenhamDxXRYD(img, x1, y1, x2, y2, colWHITE)
	}
}

func Benchmark_Bresenham(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham(img, x1, y1, x2, y2, colWHITE)
	}
}
