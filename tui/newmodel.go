package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

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

func NewResultModel(width int, height int) ResultModel {
	return ResultModel{
		width:  width,
		height: height,
		arr:    make([]int, 3),
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
