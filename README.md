# Draw line with the Bresenham algorithm in Golang

Because golang is lacking of a basic drawing library, this is how I rediscovered and implemented the bresenham algorithm to draw a line.

# Example

```go
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

	// save image in example1.png
	toimg, _ := os.Create("example1.png")
	defer toimg.Close()
	png.Encode(toimg, img)
}
```

bresenham.DrawLine() gets a Plotter interface as it's first argument.

```go
type Plotter interface {
	Set(x int, y int, c color.Color)
}
```

# Benchmark

```sh
go test -bench=.
```

NB: On a modern processor with dedicated floating point ALU, naive implementation could be on par (or faster) with the bresenham version.
