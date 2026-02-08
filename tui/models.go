package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/cubexgani/regression-calculator/utils"
)

// The choice screen
type ChoiceModel struct {
	width    int
	height   int
	opts     []string
	cursor   int
	selected bool
	isquit   bool
	input    textinput.Model
	Inswitch int
	rownum   int
	errmsg   string
}

type XYInModel struct {
	winwdth int
	winht   int

	regtype string
	n       int
	x       []float32
	y       []float32
	xytext  textinput.Model
	rowcurs int
	colcurs int
	errmsg  string
	done    bool
	rowSize int
}

type ResultModel struct {
	width  int
	height int

	cellSize int
	n        int
	regtype  string
	x        []float32
	y        []float32
	solnvec  []float32
	table    utils.Regression
	solns    []string
	errmsg   string
}

type DadModel struct {
	Choice ChoiceModel
	XYIn   XYInModel
	Result ResultModel
}
