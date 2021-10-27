package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/StephaneBunel/bresenham"
)

func main() {
	var imgRect = image.Rect(0, 0, 500, 500)
	var img = image.NewRGBA(imgRect)
	var colBLUE = color.RGBA{0, 0, 255, 255}

	// draw line
	bresenham.DrawLine(img, 14, 71, 441, 317, colBLUE)

	// save image
	toimg, _ := os.Create("example1.png")
	defer toimg.Close()
	png.Encode(toimg, img)
}
