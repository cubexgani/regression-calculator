package tui

import (
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
		width-80, height-15, minx-xAxisExtension, maxx+xAxisExtension, miny-yAxisExtension, maxy+yAxisExtension,
		linechart.WithXYSteps(4, 2),
		linechart.WithStyles(axisStyle, labelStyle, replacedStyle),
	)

	lc.DrawXYAxisAndLabel()

	switch strings.ToLower(regtype) {
	case "linear":
		aymin := vec[0] + vec[1]*float32(minx)
		aymax := vec[0] + vec[1]*float32(maxx)
		lc.DrawBrailleLineWithStyle(canvas.Float64Point{X: minx, Y: float64(aymin)}, canvas.Float64Point{X: maxx, Y: float64(aymax)}, lineStyle)

	}
	for i := range len(x) {
		lc.DrawRuneWithStyle(canvas.Float64Point{X: float64(x[i]), Y: float64(y[i])}, 'o', pointStyle)
	}
	return lc
}
