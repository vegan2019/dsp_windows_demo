package main

import (
	"github.com/Arafatk/glot"
	"github.com/mattetti/audio/dsp/windows"
)

func main() {

	var BlackmanWindow []float64
	var HammingWindow []float64
	var NuttallWindow []float64

	var nPoints []float64
	N := 128
	BlackmanWindow = windows.Blackman(N)
	HammingWindow = windows.Hamming(N)
	NuttallWindow = windows.Nuttall(N)

	for i := 0; i < N; i++ {
		nPoints = append(nPoints, float64(i))
	}

	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	style := "lines"
	plot.AddPointGroup("Blackman", style, [][]float64{nPoints, BlackmanWindow})
	plot.AddPointGroup("Hamming", style, [][]float64{nPoints, HammingWindow})
	plot.AddPointGroup("Nuttall", style, [][]float64{nPoints, NuttallWindow})

	plot.SetTitle("3 windows")
	plot.SetXLabel("samples")
	plot.SetYLabel("amplitude")
	plot.SavePlot("3 windows.jpg")
}
