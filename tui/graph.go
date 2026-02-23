package tui

import (
	"math"
	"strings"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/charmbracelet/lipgloss"
)

var replacedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("35")) // jade green

var lineStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")) // skyblue-ish

var pointStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("203")) // less striking red

var axisStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("37")) // cyan-ish

var labelStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("84")) // bright light green

func DrawGraph(regtype string, width, height int, vec, x, y []float32) linechart.Model {
	var lc linechart.Model
	var delfac float64
	delfac = 5
	var minx, miny, maxx, maxy float64
	minx, maxx, miny, maxy = float64(arrMin(x)), float64(arrMax(x)), float64(arrMin(y)), float64(arrMax(y))
	xAxisExtension := (maxx - minx) / delfac
	yAxisExtension := (maxy - miny) / delfac

	lc = linechart.New(
		width-80, height-15, math.Floor(minx-xAxisExtension), math.Ceil(maxx+xAxisExtension),
		math.Floor(miny-yAxisExtension), math.Ceil(maxy+yAxisExtension),
		// FIXME: This looks like a good enough estimate for the XY labelling in fullscreen, but I need to dynamically calculate
		// it according to the input range
		linechart.WithXYSteps(6, 3),
		linechart.WithStyles(axisStyle, labelStyle, replacedStyle),
	)

	lc.DrawXYAxisAndLabel()

	switch strings.ToLower(regtype) {
	case "linear":
		DrawLinearGraph(&lc, vec, minx, maxx)
	case "quadratic":
		DrawQuadraticGraph(&lc, vec, x, y, minx, maxx)
	}
	for i := range len(x) {
		lc.DrawRuneWithStyle(canvas.Float64Point{X: float64(x[i]), Y: float64(y[i])}, 'o', pointStyle)
	}
	return lc
}

func DrawLinearGraph(lc *linechart.Model, vec []float32, minx, maxx float64) {
	aymin := vec[0] + vec[1]*float32(minx)
	aymax := vec[0] + vec[1]*float32(maxx)
	lc.DrawBrailleLineWithStyle(canvas.Float64Point{X: minx, Y: float64(aymin)}, canvas.Float64Point{X: maxx, Y: float64(aymax)}, lineStyle)
}

func DrawQuadraticGraph(lc *linechart.Model, vec, x, y []float32, minx, maxx float64) {
	// For how granular(?) I want my quadratic curve to be
	granLevel := 10
	var p, incp float32
	f32minx, f32maxx := float32(minx), float32(maxx)
	p = f32minx
	incp = (f32maxx - f32minx) / float32(granLevel)
	aymin := vec[0] + vec[1]*f32minx + vec[2]*f32minx*f32minx

	prev := canvas.Float64Point{X: minx, Y: float64(aymin)}
	for range granLevel {
		p += incp
		yc := vec[0] + vec[1]*p + vec[2]*p*p
		cur := canvas.Float64Point{X: float64(p), Y: float64(yc)}
		lc.DrawBrailleLineWithStyle(prev, cur, lineStyle)
		prev = cur
	}
}

func arrMax(x []float32) float32 {
	max := x[0]
	for i := range len(x) {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func arrMin(x []float32) float32 {
	min := x[0]
	for i := range len(x) {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}
