package bresenham

// 2016-10-22, Stéphane Bunel

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

var imgRect = image.Rect(0, 0, 100, 100)
var img = image.NewRGBA(imgRect)
var p1 = image.Point{10, 10}
var p2 = image.Point{25, 90}
var p3 = image.Point{50, 50}
var x1, x2, y1, y2 = 4, 7, 71, 35

func drawCross(img *image.RGBA, p image.Point, col color.Color) {
	x, y := p.X, p.Y
	for s := -3; s < 4; s++ {
		img.Set(x+s, y, col)
		img.Set(x, y+s, col)
	}
}

func Test_Bresenham_1(t *testing.T) {
	drawCross(img, image.Point{imgRect.Min.X, imgRect.Min.Y}, color.White)
	drawCross(img, image.Point{imgRect.Max.X - 1, imgRect.Min.Y}, color.White)
	drawCross(img, image.Point{imgRect.Max.X - 1, imgRect.Max.Y - 1}, color.White)
	drawCross(img, image.Point{imgRect.Min.X, imgRect.Max.Y - 1}, color.White)

	drawCross(img, p1, color.White)
	drawCross(img, p2, color.White)
	drawCross(img, p3, color.White)

	Bresenham(img, p1, p2, color.RGBA{100, 100, 255, 255})
	Bresenham(img, p2, p3, color.RGBA{100, 100, 255, 255})
	Bresenham(img, p3, p1, color.RGBA{100, 100, 255, 255})

	// Save result
	toimg, _ := os.Create("drawline.png")
	defer toimg.Close()
	png.Encode(toimg, img)
	t.Log("drawline.png saved.")
}

func Benchmark_Bresenham(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham(img, p1, p2, color.RGBA{100, 255, 100, 255})
	}
}

func Benchmark_Bresenham_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_1(img, x1, y1, x2, y2, color.RGBA{100, 255, 100, 255})
	}
}

func Benchmark_Bresenham_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_1(img, x1, y1, x2, y2, color.RGBA{100, 255, 100, 255})
	}
}

func Benchmark_Bresenham_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_1(img, x1, y1, x2, y2, color.RGBA{100, 255, 100, 255})
	}
}

func Benchmark_Bresenham_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bresenham_1(img, x1, y1, x2, y2, color.RGBA{100, 255, 100, 255})
	}
}
