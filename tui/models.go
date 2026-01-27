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
	inswitch int
	rownum   int
	errmsg   string
}

type DadModel struct {
	Choice ChoiceModel
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
