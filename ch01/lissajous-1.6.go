package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x44, 0xff},
	color.RGBA{0x00, 0x00, 0x88, 0xff},
	color.RGBA{0x00, 0x00, 0xcc, 0xff},
	color.RGBA{0x00, 0x44, 0x00, 0xff},
	color.RGBA{0x00, 0x44, 0x44, 0xff},
	color.RGBA{0x00, 0x44, 0x88, 0xff},
	color.RGBA{0x00, 0x44, 0xcc, 0xff},
	color.RGBA{0x00, 0x88, 0x00, 0xff},
	color.RGBA{0x00, 0x88, 0x44, 0xff},
	color.RGBA{0x00, 0x88, 0x88, 0xff},
	color.RGBA{0x00, 0x88, 0xcc, 0xff},
	color.RGBA{0x00, 0xcc, 0x00, 0xff},
	color.RGBA{0x00, 0xcc, 0x44, 0xff},
	color.RGBA{0x00, 0xcc, 0x88, 0xff},
	color.RGBA{0x00, 0xcc, 0xcc, 0xff},
	color.RGBA{0x00, 0xdd, 0xdd, 0xff},
	color.RGBA{0x11, 0x11, 0x11, 0xff},
	color.RGBA{0x00, 0x00, 0x55, 0xff},
	color.RGBA{0x00, 0x00, 0x99, 0xff},
	color.RGBA{0x00, 0x00, 0xdd, 0xff},
	color.RGBA{0x00, 0x55, 0x00, 0xff},
	color.RGBA{0x00, 0x55, 0x55, 0xff},
	color.RGBA{0x00, 0x4c, 0x99, 0xff},
	color.RGBA{0x00, 0x49, 0xdd, 0xff},
	color.RGBA{0x00, 0x99, 0x00, 0xff},
	color.RGBA{0x00, 0x99, 0x4c, 0xff},
	color.RGBA{0x00, 0x99, 0x99, 0xff},
	color.RGBA{0x00, 0x93, 0xdd, 0xff},
	color.RGBA{0x00, 0xdd, 0x00, 0xff},
	color.RGBA{0x00, 0xdd, 0x49, 0xff},
	color.RGBA{0x00, 0xdd, 0x93, 0xff},
	color.RGBA{0x00, 0xee, 0x9e, 0xff},
	color.RGBA{0x00, 0xee, 0xee, 0xff},
	color.RGBA{0x22, 0x22, 0x22, 0xff},
	color.RGBA{0x00, 0x00, 0x66, 0xff},
	color.RGBA{0x00, 0x00, 0xaa, 0xff},
	color.RGBA{0x00, 0x00, 0xee, 0xff},
	color.RGBA{0x00, 0x66, 0x00, 0xff},
	color.RGBA{0x00, 0x66, 0x66, 0xff},
	color.RGBA{0x00, 0x55, 0xaa, 0xff},
	color.RGBA{0x00, 0x4f, 0xee, 0xff},
	color.RGBA{0x00, 0xaa, 0x00, 0xff},
	color.RGBA{0x00, 0xaa, 0x55, 0xff},
	color.RGBA{0x00, 0xaa, 0xaa, 0xff},
	color.RGBA{0x00, 0x9e, 0xee, 0xff},
	color.RGBA{0x00, 0xee, 0x00, 0xff},
	color.RGBA{0x00, 0xee, 0x4f, 0xff},
	color.RGBA{0x00, 0xff, 0x55, 0xff},
	color.RGBA{0x00, 0xff, 0xaa, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0x33, 0x33, 0x33, 0xff},
	color.RGBA{0x00, 0x00, 0x77, 0xff},
	color.RGBA{0x00, 0x00, 0xbb, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0x77, 0x00, 0xff},
	color.RGBA{0x00, 0x77, 0x77, 0xff},
	color.RGBA{0x00, 0x5d, 0xbb, 0xff},
	color.RGBA{0x00, 0x55, 0xff, 0xff},
	color.RGBA{0x00, 0xbb, 0x00, 0xff},
	color.RGBA{0x00, 0xbb, 0x5d, 0xff},
	color.RGBA{0x00, 0xbb, 0xbb, 0xff},
	color.RGBA{0x00, 0xaa, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x44, 0x00, 0x44, 0xff},
	color.RGBA{0x44, 0x00, 0x88, 0xff},
	color.RGBA{0x44, 0x00, 0xcc, 0xff},
	color.RGBA{0x44, 0x44, 0x00, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0x44, 0x44, 0x88, 0xff},
	color.RGBA{0x44, 0x44, 0xcc, 0xff},
	color.RGBA{0x44, 0x88, 0x00, 0xff},
	color.RGBA{0x44, 0x88, 0x44, 0xff},
	color.RGBA{0x44, 0x88, 0x88, 0xff},
	color.RGBA{0x44, 0x88, 0xcc, 0xff},
	color.RGBA{0x44, 0xcc, 0x00, 0xff},
	color.RGBA{0x44, 0xcc, 0x44, 0xff},
	color.RGBA{0x44, 0xcc, 0x88, 0xff},
	color.RGBA{0x44, 0xcc, 0xcc, 0xff},
	color.RGBA{0x44, 0x00, 0x00, 0xff},
	color.RGBA{0x55, 0x00, 0x00, 0xff},
	color.RGBA{0x55, 0x00, 0x55, 0xff},
	color.RGBA{0x4c, 0x00, 0x99, 0xff},
	color.RGBA{0x49, 0x00, 0xdd, 0xff},
	color.RGBA{0x55, 0x55, 0x00, 0xff},
	color.RGBA{0x55, 0x55, 0x55, 0xff},
	color.RGBA{0x4c, 0x4c, 0x99, 0xff},
	color.RGBA{0x49, 0x49, 0xdd, 0xff},
	color.RGBA{0x4c, 0x99, 0x00, 0xff},
	color.RGBA{0x4c, 0x99, 0x4c, 0xff},
	color.RGBA{0x4c, 0x99, 0x99, 0xff},
	color.RGBA{0x49, 0x93, 0xdd, 0xff},
	color.RGBA{0x49, 0xdd, 0x00, 0xff},
	color.RGBA{0x49, 0xdd, 0x49, 0xff},
	color.RGBA{0x49, 0xdd, 0x93, 0xff},
	color.RGBA{0x49, 0xdd, 0xdd, 0xff},
	color.RGBA{0x4f, 0xee, 0xee, 0xff},
	color.RGBA{0x66, 0x00, 0x00, 0xff},
	color.RGBA{0x66, 0x00, 0x66, 0xff},
	color.RGBA{0x55, 0x00, 0xaa, 0xff},
	color.RGBA{0x4f, 0x00, 0xee, 0xff},
	color.RGBA{0x66, 0x66, 0x00, 0xff},
	color.RGBA{0x66, 0x66, 0x66, 0xff},
	color.RGBA{0x55, 0x55, 0xaa, 0xff},
	color.RGBA{0x4f, 0x4f, 0xee, 0xff},
	color.RGBA{0x55, 0xaa, 0x00, 0xff},
	color.RGBA{0x55, 0xaa, 0x55, 0xff},
	color.RGBA{0x55, 0xaa, 0xaa, 0xff},
	color.RGBA{0x4f, 0x9e, 0xee, 0xff},
	color.RGBA{0x4f, 0xee, 0x00, 0xff},
	color.RGBA{0x4f, 0xee, 0x4f, 0xff},
	color.RGBA{0x4f, 0xee, 0x9e, 0xff},
	color.RGBA{0x55, 0xff, 0xaa, 0xff},
	color.RGBA{0x55, 0xff, 0xff, 0xff},
	color.RGBA{0x77, 0x00, 0x00, 0xff},
	color.RGBA{0x77, 0x00, 0x77, 0xff},
	color.RGBA{0x5d, 0x00, 0xbb, 0xff},
	color.RGBA{0x55, 0x00, 0xff, 0xff},
	color.RGBA{0x77, 0x77, 0x00, 0xff},
	color.RGBA{0x77, 0x77, 0x77, 0xff},
	color.RGBA{0x5d, 0x5d, 0xbb, 0xff},
	color.RGBA{0x55, 0x55, 0xff, 0xff},
	color.RGBA{0x5d, 0xbb, 0x00, 0xff},
	color.RGBA{0x5d, 0xbb, 0x5d, 0xff},
	color.RGBA{0x5d, 0xbb, 0xbb, 0xff},
	color.RGBA{0x55, 0xaa, 0xff, 0xff},
	color.RGBA{0x55, 0xff, 0x00, 0xff},
	color.RGBA{0x55, 0xff, 0x55, 0xff},
	color.RGBA{0x88, 0x00, 0x88, 0xff},
	color.RGBA{0x88, 0x00, 0xcc, 0xff},
	color.RGBA{0x88, 0x44, 0x00, 0xff},
	color.RGBA{0x88, 0x44, 0x44, 0xff},
	color.RGBA{0x88, 0x44, 0x88, 0xff},
	color.RGBA{0x88, 0x44, 0xcc, 0xff},
	color.RGBA{0x88, 0x88, 0x00, 0xff},
	color.RGBA{0x88, 0x88, 0x44, 0xff},
	color.RGBA{0x88, 0x88, 0x88, 0xff},
	color.RGBA{0x88, 0x88, 0xcc, 0xff},
	color.RGBA{0x88, 0xcc, 0x00, 0xff},
	color.RGBA{0x88, 0xcc, 0x44, 0xff},
	color.RGBA{0x88, 0xcc, 0x88, 0xff},
	color.RGBA{0x88, 0xcc, 0xcc, 0xff},
	color.RGBA{0x88, 0x00, 0x00, 0xff},
	color.RGBA{0x88, 0x00, 0x44, 0xff},
	color.RGBA{0x99, 0x00, 0x4c, 0xff},
	color.RGBA{0x99, 0x00, 0x99, 0xff},
	color.RGBA{0x93, 0x00, 0xdd, 0xff},
	color.RGBA{0x99, 0x4c, 0x00, 0xff},
	color.RGBA{0x99, 0x4c, 0x4c, 0xff},
	color.RGBA{0x99, 0x4c, 0x99, 0xff},
	color.RGBA{0x93, 0x49, 0xdd, 0xff},
	color.RGBA{0x99, 0x99, 0x00, 0xff},
	color.RGBA{0x99, 0x99, 0x4c, 0xff},
	color.RGBA{0x99, 0x99, 0x99, 0xff},
	color.RGBA{0x93, 0x93, 0xdd, 0xff},
	color.RGBA{0x93, 0xdd, 0x00, 0xff},
	color.RGBA{0x93, 0xdd, 0x49, 0xff},
	color.RGBA{0x93, 0xdd, 0x93, 0xff},
	color.RGBA{0x93, 0xdd, 0xdd, 0xff},
	color.RGBA{0x99, 0x00, 0x00, 0xff},
	color.RGBA{0xaa, 0x00, 0x00, 0xff},
	color.RGBA{0xaa, 0x00, 0x55, 0xff},
	color.RGBA{0xaa, 0x00, 0xaa, 0xff},
	color.RGBA{0x9e, 0x00, 0xee, 0xff},
	color.RGBA{0xaa, 0x55, 0x00, 0xff},
	color.RGBA{0xaa, 0x55, 0x55, 0xff},
	color.RGBA{0xaa, 0x55, 0xaa, 0xff},
	color.RGBA{0x9e, 0x4f, 0xee, 0xff},
	color.RGBA{0xaa, 0xaa, 0x00, 0xff},
	color.RGBA{0xaa, 0xaa, 0x55, 0xff},
	color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
	color.RGBA{0x9e, 0x9e, 0xee, 0xff},
	color.RGBA{0x9e, 0xee, 0x00, 0xff},
	color.RGBA{0x9e, 0xee, 0x4f, 0xff},
	color.RGBA{0x9e, 0xee, 0x9e, 0xff},
	color.RGBA{0x9e, 0xee, 0xee, 0xff},
	color.RGBA{0xaa, 0xff, 0xff, 0xff},
	color.RGBA{0xbb, 0x00, 0x00, 0xff},
	color.RGBA{0xbb, 0x00, 0x5d, 0xff},
	color.RGBA{0xbb, 0x00, 0xbb, 0xff},
	color.RGBA{0xaa, 0x00, 0xff, 0xff},
	color.RGBA{0xbb, 0x5d, 0x00, 0xff},
	color.RGBA{0xbb, 0x5d, 0x5d, 0xff},
	color.RGBA{0xbb, 0x5d, 0xbb, 0xff},
	color.RGBA{0xaa, 0x55, 0xff, 0xff},
	color.RGBA{0xbb, 0xbb, 0x00, 0xff},
	color.RGBA{0xbb, 0xbb, 0x5d, 0xff},
	color.RGBA{0xbb, 0xbb, 0xbb, 0xff},
	color.RGBA{0xaa, 0xaa, 0xff, 0xff},
	color.RGBA{0xaa, 0xff, 0x00, 0xff},
	color.RGBA{0xaa, 0xff, 0x55, 0xff},
	color.RGBA{0xaa, 0xff, 0xaa, 0xff},
	color.RGBA{0xcc, 0x00, 0xcc, 0xff},
	color.RGBA{0xcc, 0x44, 0x00, 0xff},
	color.RGBA{0xcc, 0x44, 0x44, 0xff},
	color.RGBA{0xcc, 0x44, 0x88, 0xff},
	color.RGBA{0xcc, 0x44, 0xcc, 0xff},
	color.RGBA{0xcc, 0x88, 0x00, 0xff},
	color.RGBA{0xcc, 0x88, 0x44, 0xff},
	color.RGBA{0xcc, 0x88, 0x88, 0xff},
	color.RGBA{0xcc, 0x88, 0xcc, 0xff},
	color.RGBA{0xcc, 0xcc, 0x00, 0xff},
	color.RGBA{0xcc, 0xcc, 0x44, 0xff},
	color.RGBA{0xcc, 0xcc, 0x88, 0xff},
	color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.RGBA{0xcc, 0x00, 0x00, 0xff},
	color.RGBA{0xcc, 0x00, 0x44, 0xff},
	color.RGBA{0xcc, 0x00, 0x88, 0xff},
	color.RGBA{0xdd, 0x00, 0x93, 0xff},
	color.RGBA{0xdd, 0x00, 0xdd, 0xff},
	color.RGBA{0xdd, 0x49, 0x00, 0xff},
	color.RGBA{0xdd, 0x49, 0x49, 0xff},
	color.RGBA{0xdd, 0x49, 0x93, 0xff},
	color.RGBA{0xdd, 0x49, 0xdd, 0xff},
	color.RGBA{0xdd, 0x93, 0x00, 0xff},
	color.RGBA{0xdd, 0x93, 0x49, 0xff},
	color.RGBA{0xdd, 0x93, 0x93, 0xff},
	color.RGBA{0xdd, 0x93, 0xdd, 0xff},
	color.RGBA{0xdd, 0xdd, 0x00, 0xff},
	color.RGBA{0xdd, 0xdd, 0x49, 0xff},
	color.RGBA{0xdd, 0xdd, 0x93, 0xff},
	color.RGBA{0xdd, 0xdd, 0xdd, 0xff},
	color.RGBA{0xdd, 0x00, 0x00, 0xff},
	color.RGBA{0xdd, 0x00, 0x49, 0xff},
	color.RGBA{0xee, 0x00, 0x4f, 0xff},
	color.RGBA{0xee, 0x00, 0x9e, 0xff},
	color.RGBA{0xee, 0x00, 0xee, 0xff},
	color.RGBA{0xee, 0x4f, 0x00, 0xff},
	color.RGBA{0xee, 0x4f, 0x4f, 0xff},
	color.RGBA{0xee, 0x4f, 0x9e, 0xff},
	color.RGBA{0xee, 0x4f, 0xee, 0xff},
	color.RGBA{0xee, 0x9e, 0x00, 0xff},
	color.RGBA{0xee, 0x9e, 0x4f, 0xff},
	color.RGBA{0xee, 0x9e, 0x9e, 0xff},
	color.RGBA{0xee, 0x9e, 0xee, 0xff},
	color.RGBA{0xee, 0xee, 0x00, 0xff},
	color.RGBA{0xee, 0xee, 0x4f, 0xff},
	color.RGBA{0xee, 0xee, 0x9e, 0xff},
	color.RGBA{0xee, 0xee, 0xee, 0xff},
	color.RGBA{0xee, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x55, 0xff},
	color.RGBA{0xff, 0x00, 0xaa, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x55, 0x00, 0xff},
	color.RGBA{0xff, 0x55, 0x55, 0xff},
	color.RGBA{0xff, 0x55, 0xaa, 0xff},
	color.RGBA{0xff, 0x55, 0xff, 0xff},
	color.RGBA{0xff, 0xaa, 0x00, 0xff},
	color.RGBA{0xff, 0xaa, 0x55, 0xff},
	color.RGBA{0xff, 0xaa, 0xaa, 0xff},
	color.RGBA{0xff, 0xaa, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x55, 0xff},
	color.RGBA{0xff, 0xff, 0xaa, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(250)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
