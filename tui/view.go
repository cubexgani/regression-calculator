package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m ChoiceModel) View() string {
	outputBuilder := strings.Builder{}
	outputBuilder.WriteString("Which regression is lwk hot?\n")
	if m.selected {
		fmt.Fprintf(&outputBuilder, "Selected: %s\n", m.opts[m.cursor])
		if m.inswitch == 1 {
			fmt.Fprintln(&outputBuilder, m.input.View())
			if m.errmsg != "" {
				fmt.Fprintln(&outputBuilder, m.errmsg)
			}
		} else {
			fmt.Fprintf(&outputBuilder, "You wrote: %s\n", m.input.Value())
			fmt.Fprintf(&outputBuilder, "Number of rows: %d\n", m.rownum)
		}
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			outputBuilder.String(),
		)
	}
	if m.isquit {
		return fmt.Sprintln("You stupid nig")
	}
	for i, opt := range m.opts {
		if i == m.cursor {
			outputBuilder.WriteString(" > ")
		} else {
			outputBuilder.WriteString("   ")
		}
		fmt.Fprintf(&outputBuilder, "%s\n", opt)
	}
	outputBuilder.WriteString("\nChoose or deth twn\n")
	fmt.Fprintln(&outputBuilder, m.input.View())
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		outputBuilder.String(),
	)
}

func (m DadModel) View() string {
	return m.Choice.View()
}
