package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m ChoiceModel) View() string {
	var (
		green       = lipgloss.Color("84")
		borderStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).
				Padding(0, 2).
				BorderForeground(green)
		bgc = lipgloss.NewStyle().
			Width(20).
			Background(lipgloss.Color("23"))
	)
	outputBuilder := strings.Builder{}
	optBuilder := strings.Builder{}
	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
	optBuilder.WriteString("Choose the type of regression\n")
	if m.selected {
		fmt.Fprint(&outputBuilder, optBuilder.String())
		fmt.Fprintf(&outputBuilder, "Selected: %s\n", m.opts[m.cursor])
		if m.Inswitch == 1 {
			fmt.Fprintln(&outputBuilder, "\nEnter number of coordinates")
			fmt.Fprintln(&outputBuilder, borderStyle.Render(m.input.View()))
			if m.errmsg != "" {
				fmt.Fprintln(&outputBuilder, m.errmsg)
			}
		}
		outputBuilder.WriteString("\n\n\n" + help.Render("[enter] Go to next screen   [ctrl+c] Exit"))
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			outputBuilder.String(),
		)
	}
	if m.isquit {
		return fmt.Sprintln("GET OUT")
	}
	for i, opt := range m.opts {
		if i == m.cursor {
			fmt.Fprintln(&optBuilder, bgc.Render("   ", opt))
		} else {
			fmt.Fprintln(&optBuilder, "   ", opt)
		}
	}
	fmt.Fprintln(&outputBuilder, borderStyle.Render(optBuilder.String()))
	outputBuilder.WriteString("\n\nEnter number of coordinates\n")
	fmt.Fprintln(&outputBuilder, m.input.View())
	outputBuilder.WriteString("\n\n\n" + help.Render("[↑/↓] Navigate    [enter] Go to number input   [ctrl+c] Exit"))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		outputBuilder.String(),
	)
}

// My dumbass drew all of this by hand. Might consider using more lipgloss borders later on.

func (m XYInModel) View() string {
	sb := strings.Builder{}
	fmt.Fprintln(&sb, "Enter the coordinates")
	fmt.Fprintln(&sb)
	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
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
		if m.rowSize == 0 {
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
			Width(m.rowSize).
			Align(lipgloss.Center)

		bgc := ls.
			Background(lipgloss.Color("23"))

		bdc := lipgloss.NewStyle().
			Foreground(lipgloss.Color("84"))

		sb.WriteString("\n")
		// top border
		sb.WriteString(bdc.Render("┌"))
		for range m.rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┬"))
		for range m.rowSize {
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
		for range m.rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┼"))
		for range m.rowSize {
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
		for range m.rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┴"))
		for range m.rowSize {
			sb.WriteString(bdc.Render("─"))
		}
		sb.WriteString(bdc.Render("┘\n"))
	}
	if m.errmsg != "" {
		sb.WriteString(m.errmsg)
		sb.WriteRune('\n')
	}
	sb.WriteString("\n\n\n")
	fmt.Fprintln(&sb, help.Render("[tab] Go to next cell     [shift+tab] Go to previous cell"))
	fmt.Fprintln(&sb, help.Render("[enter] Confirm inputs    [ctrl + c] Exit"))
	return lipgloss.Place(
		m.winwdth,
		m.winht,
		lipgloss.Center,
		lipgloss.Center,
		sb.String(),
	)
}

// TODO: A couple of things to do here:
// - Include tabs
// - Optionally show formulae for the system of equations
// - Include viewport for vertical scrolling
// On second thought, I'm ditching viewports for now, looks like I gotta make some not-so-fun changes. Or I'm overthinking,
// whatever.

func (m ResultModel) View() string {
	var (
		green = lipgloss.Color("84")
		gray  = lipgloss.Color("245")

		headerStyle = lipgloss.NewStyle().Foreground(green).Bold(true).Align(lipgloss.Center)
		// TODO: make cell width dependent on screen size and number of columns
		cellStyle   = lipgloss.NewStyle().Padding(0, 1).Width(m.cellSize)
		rowStyle    = cellStyle.Foreground(gray)
		borderStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).
				BorderForeground(green)
		graphBorderStyle = borderStyle.Padding(2, 4)
		helpStyle        = lipgloss.NewStyle().
					Foreground(lipgloss.Color("241"))
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(green)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case table.HeaderRow:
				return headerStyle
			default:
				return rowStyle
			}
		})

	// For printing data in a completely raw format, using tabs. Keeping it for debugging purposes
	outputBuilder := strings.Builder{}
	// For the beautiful lipgloss table
	tableBuilder := strings.Builder{}

	fmt.Fprintln(&tableBuilder, "Regression Result")
	fmt.Fprintln(&tableBuilder)

	if m.graphMode && m.graphInitted {
		fmt.Fprintln(&tableBuilder, graphBorderStyle.Render(m.lc.View()))
		fmt.Fprintln(&tableBuilder)
		fmt.Fprintln(&tableBuilder, helpStyle.Render("[r] Go to table view    [ctrl + c] Exit"))
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			tableBuilder.String(),
		)
	}

	if m.table != nil {
		// For now, I hardcoded these sizes.
		if m.cellSize == 0 {
			return lipgloss.Place(
				m.width,
				m.height,
				lipgloss.Center,
				lipgloss.Center,
				"WHEN WIL'ST THOU WIDEN THY SCREEN?",
			)
		}
		xv, yv := m.table.GetData()
		// still debating on whether I should use the len function or the exported variables
		// Powers_Xi and Powers_YXi
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
		for i := range xl {
			ele := strconv.FormatFloat(float64(xv.Sums[i]), 'f', 2, 32)
			if i == 0 {
				arx[i] = fmt.Sprintf("Σx = %s", ele)
			} else {
				arx[i] = fmt.Sprintf("Σx^%d = %s", i+1, ele)
			}
		}
		for i := range yl {
			ele := strconv.FormatFloat(float64(yv.Sums[i]), 'f', 2, 32)
			switch i {
			case 0:
				ary[i] = fmt.Sprintf("Σy = %s", ele)
			case 1:
				ary[i] = fmt.Sprintf("Σyx = %s", ele)

			default:
				ary[i] = fmt.Sprintf("Σyx^%d = %s", i, ele)
			}
		}
		fmt.Fprintln(&outputBuilder, xv.Sums, yv.Sums)
		t.Row(append(arx, ary...)...)
		fmt.Fprintln(&tableBuilder, t.String())
		fmt.Fprintln(&tableBuilder, "\nSystem of equations to solve:")
		fmt.Fprintln(&tableBuilder, borderStyle.Render(getEqnSystem(m.regtype, xv.Sums, yv.Sums, m.n)))
		if m.solnvec != nil {
			fmt.Fprintln(&tableBuilder, "\nSoltion vector:")
			fmt.Fprintln(&tableBuilder, borderStyle.Render(getVec(m.regtype, m.solnvec)))
		}
		fmt.Fprintln(&tableBuilder, "\nEquation of curve:")
		fmt.Fprintln(&tableBuilder, borderStyle.Render(m.solns[0]))
	}

	fmt.Fprintf(&outputBuilder, "%v\n", m.solns)
	fmt.Fprintf(&outputBuilder, "%v\n", m.errmsg)
	fmt.Fprintln(&tableBuilder, m.errmsg)
	for range 3 {
		fmt.Fprintln(&tableBuilder)
	}
	fmt.Fprintln(&tableBuilder, helpStyle.Render("[r] Go to curve view    [ctrl + c] Exit"))
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		// outputBuilder.String(),
		tableBuilder.String(),
	)
}

// TODO: Softcode this ig?

func getEqnSystem(regtype string, xsums []float32, ysums []float32, n int) string {
	eqnsBuilder := strings.Builder{}
	switch strings.ToLower(regtype) {
	case "linear":
		fmt.Fprintf(&eqnsBuilder, "%.2f = %da + %.2fb\n"+
			"%.2f = %.2fa + %.2fb", ysums[0],
			n, xsums[0], ysums[1], xsums[0], xsums[1],
		)
	case "quadratic":
		fmt.Fprintf(&eqnsBuilder, "%.2f = %da + %.2fb + %.2fc\n"+
			"%.2f = %.2fa + %.2fb + %.2fc\n"+
			"%.2f = %.2fa + %.2fb + %.2fc",
			ysums[0], n, xsums[0], xsums[1], ysums[1], xsums[0], xsums[1], xsums[2],
			ysums[2], xsums[1], xsums[2], xsums[3],
		)
	}
	return eqnsBuilder.String()
}

func getVec(regtype string, solnvec []float32) string {
	vecBuilder := strings.Builder{}
	switch strings.ToLower(regtype) {
	case "linear":
		fmt.Fprintf(&vecBuilder, "[a b] = %v", solnvec)
	case "quadratic":
		fmt.Fprintf(&vecBuilder, "[a b c] = %v", solnvec)
	}
	return vecBuilder.String()
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
