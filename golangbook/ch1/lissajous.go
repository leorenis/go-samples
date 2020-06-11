// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{color.RGBA{0x00, 0x7d, 0x9c, 0xff}, color.RGBA{0x00, 0x00, 0x00, 0xff}}

const (
	whiteIndex = 0 // first pallet's color.
	greenIndex = 1 // next pallet's color.
)

// ShowLissajous is
func ShowLissajous() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout, 20)
}

// To run this program type: `go run book.go > out.gif`. Remember: In book.go you have to call ch1.ShowLissajous()
// cycles: número de revolucoes completas do oscilador x
func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 // resolucao angular
		size    = 100   // canvas da imagem cobre de [-size..+size]
		nframes = 64    // número de quadros da animacao
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0 // frequencia relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferenca de fase

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificacao
}
