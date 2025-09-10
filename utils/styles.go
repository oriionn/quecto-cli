package utils

import "github.com/charmbracelet/lipgloss"

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF0000"))

var SuccessStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#00FF00"))
