package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m ChoiceModel) View() string {
	outputBuilder := strings.Builder{}
	outputBuilder.WriteString("Which regression is lwk hot?\n")
	if m.selected {
		fmt.Fprintf(&outputBuilder, "Selected: %s\n", m.opts[m.cursor])
		if m.Inswitch == 1 {
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

func (m XYInModel) View() string {
	sb := strings.Builder{}
	if m.done {
		sb.WriteString("x, y\n")
		for i := range m.n {
			sb.WriteString(strconv.FormatFloat(float64(m.x[i]), 'f', -1, 32))
			sb.WriteString(", ")
			sb.WriteString(strconv.FormatFloat(float64(m.y[i]), 'f', -1, 32))

			sb.WriteRune('\n')
		}
		return sb.String()

	} else {
		if rowSize == 0 {
			return lipgloss.Place(
				m.winwdth,
				m.winht,
				lipgloss.Center,
				lipgloss.Center,
				"OH GOOD HEAVENS! WIDEN THY SCREEN!\n",
			)

		}
		// styles: ls for alignment, bgc for bg color of text boxes, bdc for border colour
		ls := lipgloss.NewStyle().
			Width(rowSize).
			Align(lipgloss.Center)

		bgc := ls.
			Background(lipgloss.Color("#21412d"))

		bdc := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#05e66e"))

		sb.WriteString("\n")
		// top border
		sb.WriteString(bdc.Render("┌"))
		for range rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┬"))
		for range rowSize {
			sb.WriteString(bdc.Render("─"))
		}

		// I have to render the newlines separately otherwise a weird space gets prepended at the beginning
		sb.WriteString(bdc.Render("┐"))
		sb.WriteString("\n")

		// heading
		sb.WriteString(bdc.Render("│"))
		sb.WriteString(ls.Render("x"))
		sb.WriteString(bdc.Render("│"))
		sb.WriteString(ls.Render("y"))
		sb.WriteString(bdc.Render("│"))
		sb.WriteString("\n")
		sb.WriteString(bdc.Render("├"))
		for range rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┼"))
		for range rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┤"))

		// textboxes
		for i := range m.n {
			var renderStyle lipgloss.Style
			if m.xytext[i][0].Focused() {
				renderStyle = bgc
			} else {
				renderStyle = ls
			}
			sb.WriteString("\n")
			sb.WriteString(bdc.Render("│"))

			sb.WriteString(renderStyle.Render(m.xytext[i][0].View()))
			sb.WriteString(bdc.Render("│"))

			if m.xytext[i][1].Focused() {
				renderStyle = bgc
			} else {
				renderStyle = ls
			}
			sb.WriteString(renderStyle.Render(m.xytext[i][1].View()))
			sb.WriteString(bdc.Render("│"))
		}
		sb.WriteString("\n")
		sb.WriteString(bdc.Render("└"))
		for range rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┴"))
		for range rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┘\n"))
	}
	if m.errmsg != "" {
		sb.WriteString(m.errmsg)
		sb.WriteRune('\n')
	}
	return lipgloss.Place(
		m.winwdth,
		m.winht,
		lipgloss.Center,
		lipgloss.Center,
		sb.String(),
	)
}

func (m ResultModel) View() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		"WELCOME TO FRANCE",
	)
}

func (m DadModel) View() string {
	if m.Choice.Inswitch > 1 && m.XYIn.done {
		return m.Result.View()
	} else if m.Choice.Inswitch > 1 {
		return m.XYIn.View()
	} else {
		return m.Choice.View()
	}
}
