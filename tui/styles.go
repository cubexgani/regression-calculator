package tui

import "github.com/charmbracelet/lipgloss"

var blueStone = lipgloss.Color("23")
var green = lipgloss.Color("84")
var gray = lipgloss.Color("245")

var replacedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("35")) // jade green

var lineStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")) // skyblue-ish

var pointStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("203")) // less striking red

var errorStyle = pointStyle

var axisStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("37")) // cyan-ish

var labelStyle = lipgloss.NewStyle().
	Foreground(green) // bright light green

var helpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241"))
