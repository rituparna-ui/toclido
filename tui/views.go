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
		render = lipgloss.JoinVertical(
			lipgloss.Center,
			centerText,
			lipgloss.
				NewStyle().
				Width(m.Window.Width).
				Render(
					m.EntryView.spinner.View()+" Loading...",
				),
		)
	}

	return centerDisplay.Render(render)
}

func BuildHomeScreen(m *RootModel) string {
	str := ""
	for _, todo := range m.HomeView.Todos {
		str += lipgloss.JoinHorizontal(lipgloss.Center, todo.Title, " : ", todo.Description)
		str += "\n"
	}
	return str
}

func BuildErrorScreen(m *RootModel) string {
	centerText := lipgloss.
		NewStyle().
		Width(m.Window.Width).
		Height(m.Window.Height).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	text := lipgloss.
		NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Foreground(lipgloss.Color("#FF0000")).
		Render(lipgloss.JoinVertical(lipgloss.Center, "Some Error Occured !!!", "Please retry", "Press Any Key To Exit"))

	return centerText.Render(text)
}
