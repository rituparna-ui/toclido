package main

import "github.com/charmbracelet/lipgloss"

var (
	EntryTextStyle = lipgloss.
			NewStyle().
			Bold(true).
			AlignHorizontal(lipgloss.Center).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("#FF00FF")).
			BorderForeground(lipgloss.Color("62"))

	EntryBlinkStyle = lipgloss.
			NewStyle().
			Blink(true).
			Foreground(lipgloss.Color("#008080"))
)
