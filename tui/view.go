package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m ChoiceModel) View() string {
	output := "Which regression is lwk hot?\n"
	if m.selected {
		output = fmt.Sprintf("%sSelected: %s\n", output, m.opts[m.cursor])
		if m.inswitch == 1 {
			output = fmt.Sprintln(output, m.input.View())
		} else {
			output = fmt.Sprintf("%sYou wrote: %s\n", output, m.input.Value())
		}
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			output,
		)
	}
	if m.isquit {
		return fmt.Sprintln("You stupid nig")
	}
	for i, opt := range m.opts {
		if i == m.cursor {
			output = fmt.Sprintf("%s > ", output)
		} else {
			output = fmt.Sprintf("%s   ", output)
		}
		output = fmt.Sprintf("%s%s\n", output, opt)
	}
	output = fmt.Sprintln(output, "\nChoose or deth twn")
	output = fmt.Sprintln(output, m.input.View())
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		output,
	)
}
