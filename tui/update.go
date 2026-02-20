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
			m.rowSize = 0
		} else if m.winwdth < 90 {
			m.rowSize = 25
		} else if m.winwdth < 140 {
			m.rowSize = 40
		} else {
			m.rowSize = 60
		}
	case tea.KeyMsg:
		switch mtype.String() {
		case "ctrl+c":
			m.xytext.Blur()
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
			if m.colcurs == 0 {
				if m.x[m.rowcurs] != 0 {
					m.xytext.SetValue(strconv.FormatFloat(float64(m.x[m.rowcurs]), 'f', -1, 32))
				} else {
					m.xytext.SetValue("")
				}
			} else {
				if m.y[m.rowcurs] != 0 {
					m.xytext.SetValue(strconv.FormatFloat(float64(m.y[m.rowcurs]), 'f', -1, 32))

				} else {
					m.xytext.SetValue("")
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
			if m.colcurs == 0 {
				if m.x[m.rowcurs] != 0 {
					m.xytext.SetValue(strconv.FormatFloat(float64(m.x[m.rowcurs]), 'f', -1, 32))
				} else {
					m.xytext.SetValue("")
				}
			} else {
				if m.y[m.rowcurs] != 0 {
					m.xytext.SetValue(strconv.FormatFloat(float64(m.y[m.rowcurs]), 'f', -1, 32))

				} else {
					m.xytext.SetValue("")
				}
			}
			fcmds = m.blurRest()

		case "enter":
			// value at current cursor is read
			ival, err := m.getVal()
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
			m.xytext.Blur()
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

func (m XYInModel) getVal() (float32, error) {
	content := m.xytext.Value()
	if content == "" {
		return 0, nil
	}
	val, err := strconv.ParseFloat(content, 32)
	return float32(val), err
}

func (m *XYInModel) updateInputs(msg tea.Msg) []tea.Cmd {
	changedModel, cmd := m.xytext.Update(msg)
	m.xytext = changedModel
	return []tea.Cmd{cmd}
}

func (m *XYInModel) blurRest() []tea.Cmd {
	return nil

}

func (m ResultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Well for now, I'm hardcoding the window width
		if m.width < 115 {
			m.cellSize = 0
		} else {
			m.cellSize = 14
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "r":
			m.graphMode = !m.graphMode
		}
	}
	return m, nil
}

func (m DadModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var changedModel tea.Model
	if m.Choice.Inswitch > 1 && m.XYIn.done {
		if m.Result.n == 0 {
			m.Result = NewResultModel(m.XYIn.winwdth, m.XYIn.winht, m.XYIn.n, m.XYIn.x, m.XYIn.y, m.XYIn.regtype)
		}
		changedModel, cmd = m.Result.Update(msg)
		m.Result = changedModel.(ResultModel)
	} else if m.Choice.Inswitch > 1 {
		if m.XYIn.x == nil {
			m.XYIn = NewXYModel(m.Choice.rownum, m.Choice.width, m.Choice.height, m.Choice.opts[m.Choice.cursor])
		}
		changedModel, cmd = m.XYIn.Update(msg)
		m.XYIn = changedModel.(XYInModel)
	} else {
		changedModel, cmd = m.Choice.Update(msg)
		m.Choice = changedModel.(ChoiceModel)
	}
	return m, cmd
}
