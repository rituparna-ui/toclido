package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func BuildWelcomeScreen(m *RootModel) string {
	centerText := lipgloss.JoinVertical(
		lipgloss.Center,
		"Welcome back to toclido 👋, "+strings.Split(m.Auth.Name, " ")[0]+" !!!",
		"Access all your todos right from the terminal",
		"Logged in as: "+m.Auth.Email,
	)

	centerText = EntryTextStyle.Render(centerText)
	blinkText := EntryBlinkStyle.Width(m.Window.Width).Render("Press Enter to continue...")

	centerDisplay := lipgloss.
		NewStyle().
		Width(m.Window.Width).
		Height(m.Window.Height).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	render := lipgloss.JoinVertical(lipgloss.Center, centerText, blinkText)

	if m.EntryView.showLoader {
		render = lipgloss.JoinVertical(lipgloss.Center, centerText, m.EntryView.progress.View())
	}

	return centerDisplay.Render(render)
}

func BuildHomeScreen(m *RootModel) string {
	return "Home Screen"
}
