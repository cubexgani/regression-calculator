package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cubexgani/regression-calculator/utils"
)

func NewChoiceModel() ChoiceModel {
	ti := textinput.New()
	ti.Prompt = "|* "
	ti.Width = 30
	return ChoiceModel{
		opts: []string{
			"Linear",
			"Logarithmic",
			"Power",
			"Exponential",
			"Quadratic",
		},
		input: ti,
	}
}

func NewXYModel(rows int, width int, height int, regtype string) XYInModel {
	num := rows

	xyt := textinput.New()
	xyt.Prompt = ""
	// rev := Reverse(regtype)
	xyt.Focus()
	return XYInModel{
		winwdth: width,
		winht:   height,
		regtype: regtype,
		n:       num,
		x:       make([]float32, num),
		y:       make([]float32, num),
		xytext:  xyt,
		rowSize: 60,
	}
}

func NewResultModel(width int, height int, n int, x, y []float32, regtype string) ResultModel {
	var sln []string
	var err error
	var tb utils.Regression
	var vec []float32

	tb, err = utils.InitTable(n, x, y, regtype)
	em := ""
	if err != nil {
		em = err.Error()
	} else {
		vec, sln, err = tb.Solve()
		if err != nil {
			em = err.Error()
		}
	}

	return ResultModel{
		n:        n,
		cellSize: 14,
		regtype:  regtype,
		width:    width,
		height:   height,
		x:        x,
		y:        y,
		table:    tb,
		solns:    sln,
		solnvec:  vec,
		errmsg:   em,
	}
}

func (m ChoiceModel) Init() tea.Cmd {
	return nil
}

func (m DadModel) Init() tea.Cmd {
	return nil
}

func (m XYInModel) Init() tea.Cmd {
	return nil
}
func (m ResultModel) Init() tea.Cmd {
	return nil
}
