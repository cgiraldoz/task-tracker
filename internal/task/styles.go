package task

import (
	"github.com/charmbracelet/lipgloss"
)

func NewTaskStyle(color Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Bold(true)
}
