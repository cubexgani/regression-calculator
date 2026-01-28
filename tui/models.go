package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
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
	n       int
	x       []int
	y       []int
	xytext  [][]textinput.Model
	rowcurs int
	colcurs int
	errmsg  string
	done    bool
	rowSize int
}

type ResultModel struct {
	width  int
	height int
	errmsg string
	arr    []int
}

var rowSize int

type DadModel struct {
	Choice ChoiceModel
	XYIn   XYInModel
	Result ResultModel
}
