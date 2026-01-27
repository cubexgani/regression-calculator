package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ChoiceModel struct {
	width    int
	height   int
	opts     []string
	cursor   int
	selected bool
	isquit   bool
	input    textinput.Model
	inswitch int
}

func NewModel() ChoiceModel {
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
