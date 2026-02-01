package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
			Foreground(lipgloss.Color("84"))

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

		// heading
		// i have to render the newlines separately otherwise a weird space gets prepended at the beginning
		sb.WriteString(bdc.Render("┐"))
		sb.WriteString("\n")
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
			sb.WriteString("\n")
			sb.WriteString(bdc.Render("│"))

			if m.rowcurs == i && m.colcurs == 0 {
				sb.WriteString(bgc.Render(m.xytext.View()))
			} else {
				sb.WriteString(ls.Render(strconv.FormatFloat(float64(m.x[i]), 'f', -1, 32)))

			}

			sb.WriteString(bdc.Render("│"))

			if m.rowcurs == i && m.colcurs == 1 {
				sb.WriteString(bgc.Render(m.xytext.View()))
			} else {
				sb.WriteString(ls.Render(strconv.FormatFloat(float64(m.y[i]), 'f', -1, 32)))
			}
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
	var (
		green = lipgloss.Color("84")
		gray  = lipgloss.Color("245")

		headerStyle = lipgloss.NewStyle().Foreground(green).Bold(true).Align(lipgloss.Center)
		cellStyle   = lipgloss.NewStyle().Padding(0, 1).Width(14)
		oddRowStyle = cellStyle.Foreground(gray)
		// evenRowStyle = cellStyle.Foreground(lightGray)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(green)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case table.HeaderRow:
				return headerStyle
			default:
				return oddRowStyle
			}
		})

	// For printing data in a completely raw format, using tabs. Keeping it for debugging purposes
	outputBuilder := strings.Builder{}
	// For the beautiful lipgloss table
	tableBuilder := strings.Builder{}

	if m.table != nil {
		xv, yv := m.table.GetData()
		// still debating on whether I should use the len function or the exported variables
		// Powers_Xi and Powers_Yi
		xl, yl := len(xv.Powers), len(yv.Powers)
		arx := make([]string, xl)
		fmt.Fprint(&outputBuilder, "x")
		arx[0] = "x"
		for i := range xl - 1 {
			fmt.Fprintf(&outputBuilder, "\tx^%d", i+2)
			arx[i+1] = fmt.Sprintf("x^%d", i+2)
		}
		ary := make([]string, yl)
		fmt.Fprint(&outputBuilder, "\t")
		fmt.Fprint(&outputBuilder, "y")
		ary[0] = "y"
		for i := range yl - 1 {
			if i == 0 {
				fmt.Fprint(&outputBuilder, "\tyx")
				ary[i+1] = "yx"
				continue
			}
			fmt.Fprintf(&outputBuilder, "\txy^%d", i+1)
			ary[i+1] = fmt.Sprintf("yx^%d", i+1)
		}
		t.Headers(append(arx, ary...)...)
		fmt.Fprintln(&outputBuilder)
		for i := range m.n {
			for j := range xl {
				ele := strconv.FormatFloat(float64(xv.Powers[j][i]), 'f', -1, 32)
				outputBuilder.WriteString(ele)
				outputBuilder.WriteString("\t")
				arx[j] = ele
			}
			for j := range yl {
				ele := strconv.FormatFloat(float64(yv.Powers[j][i]), 'f', -1, 32)
				outputBuilder.WriteString(ele)
				outputBuilder.WriteString("\t")
				ary[j] = ele
			}
			t.Row(append(arx, ary...)...)
			outputBuilder.WriteString("\n")
		}
		tableBuilder.WriteString(t.String())
		fmt.Fprintln(&tableBuilder)
		fmt.Fprintln(&tableBuilder, m.solns)
	}

	fmt.Fprintf(&outputBuilder, "%v\n", m.solns)
	fmt.Fprintf(&outputBuilder, "%v\n", m.errmsg)
	fmt.Fprintln(&tableBuilder, m.errmsg)
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		// outputBuilder.String(),
		tableBuilder.String(),
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
