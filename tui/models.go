package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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

var rowSize int

type DadModel struct {
	Choice ChoiceModel
	XYIn   XYInModel
}

func NewChoiceModel() ChoiceModel {
	ti := textinput.New()
	ti.Prompt = "|* "
	return ChoiceModel{
		opts: []string{
			"Linear",
			"Logarithmic",
			"Power",
			"Exponential",
		},
		input: ti,
	}
}

func (m ChoiceModel) Init() tea.Cmd {
	return nil
}

func (m DadModel) Init() tea.Cmd {
	return nil
}

func NewXYModel(rows int, width int, height int) XYInModel {
	rowSize = 60
	num := rows

	xyt := make([][]textinput.Model, num)
	for i := range num {
		xyt[i] = make([]textinput.Model, 2)
		for j := range 2 {
			xyt[i][j] = textinput.New()
			xyt[i][j].Prompt = ""
			xyt[i][j].Placeholder = "0"
		}
	}
	xyt[0][0].Focus()
	return XYInModel{
		winwdth: width,
		winht:   height,
		n:       num,
		x:       make([]int, num),
		y:       make([]int, num),
		xytext:  xyt,
	}
}

func (m XYInModel) Init() tea.Cmd {
	return nil
}
