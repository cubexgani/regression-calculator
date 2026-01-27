package tui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m ChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmd = nil
	switch mtype := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = mtype.Width
		m.height = mtype.Height
	case tea.KeyMsg:
		switch mtype.String() {
		case "ctrl+c", "q":
			m.isquit = true
			m.input.Blur()
			return m, tea.Quit
		case "down":
			if m.inswitch == 0 {
				m.cursor++
				if m.cursor == len(m.opts) {
					m.cursor = 0
				}
			}
		case "up":
			if m.inswitch == 0 {
				m.cursor--
				if m.cursor < 0 {
					m.cursor = len(m.opts) - 1
				}
			}
		case "enter":
			m.selected = true
			m.inswitch++
			if m.inswitch > 1 {
				st := m.input.Value()
				val, err := strconv.ParseInt(st, 10, 32)
				if err != nil {
					m.inswitch--
					m.errmsg = err.Error()
				} else {
					m.rownum = int(val)
					if m.rownum < 2 {
						m.inswitch--
						m.errmsg = "Please enter a number more than 1"
					}
				}
				if m.inswitch > 1 {
					m.errmsg = ""
				}
				return m, nil
			}
		}
	}
	if m.inswitch == 1 {
		var upcmd tea.Cmd
		m.input, upcmd = m.input.Update(msg)
		// cmd = upcmd
		cmd = tea.Batch(upcmd, m.input.Focus())
	} else {
		m.input.Blur()
	}
	return m, cmd
}

func (m DadModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.Choice.Update(msg)
}
