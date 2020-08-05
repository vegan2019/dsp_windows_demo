package main

import (
	"fmt"

	"github.com/Arafatk/glot"
	"github.com/mattetti/audio/dsp/windows"
)

func HammingWindowPlot() {
	var HammingWindow []float64
	var tPoints []float64

	N := 128
	HammingWindow = windows.Hamming(N)

	// horizontal axis for plot
	for i := 0; i < N; i++ {
		tPoints = append(tPoints, float64(i))
	}

	PlotXY(tPoints, HammingWindow, "Hamming Window Plot", "HammingWindowPlot.jpg")

}

func PlotXY(xPoints, yPoints []float64, title string, figName string) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	if len(yPoints) != len(xPoints) {
		fmt.Println("len(yPoints) != len(xPoints)")
		return
	}

	points := [][]float64{xPoints, yPoints}

	// Adding a point group
	pointGroupName := title
	style := "lines"
	plot.AddPointGroup(pointGroupName, style, points)
	plot.SetTitle(title)
	plot.SetXLabel("Samples")
	plot.SetYLabel("Amplitude")
	plot.SavePlot(figName)
}

func main() {
	HammingWindowPlot()
}
