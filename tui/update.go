package tui

import (
	"fmt"
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
		case "ctrl+c":
			m.isquit = true
			m.input.Blur()
			return m, tea.Quit
		case "down":
			if m.Inswitch == 0 {
				m.cursor++
				if m.cursor == len(m.opts) {
					m.cursor = 0
				}
			}
		case "up":
			if m.Inswitch == 0 {
				m.cursor--
				if m.cursor < 0 {
					m.cursor = len(m.opts) - 1
				}
			}
		case "enter":
			m.selected = true
			m.Inswitch++
			if m.Inswitch > 1 {
				st := m.input.Value()
				val, err := strconv.ParseInt(st, 10, 32)
				if err != nil {
					m.Inswitch--
					m.errmsg = err.Error()
				} else {
					m.rownum = int(val)
					if m.rownum < 2 {
						m.Inswitch--
						m.errmsg = "Please enter a number more than 1"
					}
				}
				if m.Inswitch > 1 {
					m.errmsg = ""
				}
				return m, nil
			}
		}
	}
	if m.Inswitch == 1 {
		var upcmd tea.Cmd
		m.input, upcmd = m.input.Update(msg)
		// cmd = upcmd
		cmd = tea.Batch(upcmd, m.input.Focus())
	} else {
		m.input.Blur()
	}
	return m, cmd
}

func (m XYInModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var focusCmd tea.Cmd
	fcmds := make([]tea.Cmd, 10)
	switch mtype := msg.(type) {
	case tea.WindowSizeMsg:
		m.winwdth = mtype.Width
		m.winht = mtype.Height

		if m.winwdth < 70 {
			rowSize = 0
		} else if m.winwdth < 90 {
			rowSize = 25
		} else if m.winwdth < 140 {
			rowSize = 40
		} else {
			rowSize = 60
		}
	case tea.KeyMsg:
		switch mtype.String() {
		case "ctrl+c":
			for i := range m.n {
				for j := range 2 {
					m.xytext[i][j].Blur()
				}
			}
			return m, tea.Quit
		case "tab":
			val, err := m.getVal()
			if err != nil {
				m.errmsg = fmt.Sprintf("%s at (%d, %d)", err, m.rowcurs, m.colcurs)
				fcmds = append(fcmds, m.blurRest()...)
				fcmds = append(fcmds, m.updateInputs(msg)...)
				return m, tea.Batch(fcmds...)
			}
			if m.colcurs == 0 {
				m.x[m.rowcurs] = val
			} else {
				m.y[m.rowcurs] = val
			}

			m.colcurs++
			if m.colcurs > 1 {
				m.colcurs = 0
				m.rowcurs++
				if m.rowcurs >= m.n {
					m.rowcurs = 0
				}
			}
			fcmds = m.blurRest()

		case "shift+tab":
			val, err := m.getVal()
			if err != nil {
				m.errmsg = fmt.Sprintf("%s at (%d, %d)", err, m.rowcurs, m.colcurs)
				fcmds = append(fcmds, m.blurRest()...)
				fcmds = append(fcmds, m.updateInputs(msg)...)
				return m, tea.Batch(fcmds...)
			}
			if m.colcurs == 0 {
				m.x[m.rowcurs] = val
			} else {
				m.y[m.rowcurs] = val
			}

			m.colcurs--
			if m.colcurs < 0 {
				m.colcurs = 1
				m.rowcurs--
				if m.rowcurs < 0 {
					m.rowcurs = m.n - 1
				}
			}
			fcmds = m.blurRest()

		case "enter":
			// value at current cursor is read
			ival, err := strconv.Atoi(m.xytext[m.rowcurs][m.colcurs].Value())
			if err != nil {
				m.errmsg = fmt.Sprintf("%s at (%d, %d)", err, m.rowcurs, m.colcurs)
				fcmds = append(fcmds, m.blurRest()...)
				fcmds = append(fcmds, m.updateInputs(msg)...)
				return m, tea.Batch(fcmds...)
			}
			if m.colcurs == 0 {
				m.x[m.rowcurs] = ival
			} else {
				m.y[m.rowcurs] = ival
			}
			for i := range m.n {
				for j := range 2 {
					m.xytext[i][j].Blur()
				}
			}
			m.errmsg = ""
			m.done = true
			return m, nil
		}
	}
	upcmds := m.updateInputs(msg)
	fcmds = append(fcmds, upcmds...)
	focusCmd = tea.Batch(fcmds...)
	return m, focusCmd
}

func (m XYInModel) getVal() (int, error) {
	content := m.xytext[m.rowcurs][m.colcurs].Value()
	if content == "" {
		return 0, nil
	}
	val, err := strconv.ParseInt(content, 10, 32)
	return int(val), err
}

func (m *XYInModel) updateInputs(msg tea.Msg) []tea.Cmd {
	cmds := make([]tea.Cmd, m.n*2)

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.n {
		for j := range 2 {
			m.xytext[i][j], cmds[2*i+j] = m.xytext[i][j].Update(msg)
		}
	}

	return cmds
}

func (m *XYInModel) blurRest() []tea.Cmd {
	fcmds := make([]tea.Cmd, m.n*2)
	for i := range m.n {
		for j := range 2 {
			if i == m.rowcurs && j == m.colcurs {
				fcmds[2*i+j] = m.xytext[i][j].Focus()
				continue
			}
			m.xytext[i][j].Blur()
		}
	}
	return fcmds

}

func (m DadModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var changedModel tea.Model
	cmds := make([]tea.Cmd, 2)
	cmdInd := 0
	if m.Choice.Inswitch > 1 {
		if m.XYIn.x == nil {
			m.XYIn = NewXYModel(m.Choice.rownum, m.Choice.width, m.Choice.height)
			cmds[cmdInd] = tea.ClearScreen
			cmdInd++
		}
		changedModel, cmd = m.XYIn.Update(msg)
		m.XYIn = changedModel.(XYInModel)
		cmds[cmdInd] = cmd
		cmdInd++
	} else {
		changedModel, cmd = m.Choice.Update(msg)
		m.Choice = changedModel.(ChoiceModel)
		cmds[cmdInd] = cmd
		cmdInd++
	}
	if cmdInd > 1 {
		return m, tea.Batch(cmds...)
	} else {
		return m, cmd
	}
}
