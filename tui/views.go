package main

import (
	"fmt"
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

func BuildTodosList(m *RootModel) string {
	availableHeight := m.Window.Height - 2

	fullScreen := lipgloss.NewStyle().
		Padding(1, 1).
		Width(m.Window.Width).
		Height(availableHeight)

	todos := ""

	for i, todo := range m.HomeView.Todos {
		x := " "
		if todo.Completed {
			x = "x"
		}
		todoItem := fmt.Sprintf("[%s] %s\n    %s", x, todo.Title, todo.Description)
		isActive := m.HomeView.index == i

		leftPad := 1
		if isActive {
			leftPad = 0
		}

		todos += lipgloss.
			NewStyle().
			Width(m.Window.Width-4).
			Border(lipgloss.RoundedBorder(), isActive).
			BorderForeground(lipgloss.Color("#008080")).
			Padding(leftPad, 0, leftPad, leftPad).
			Render(todoItem)
		todos += "\n"
	}

	return fullScreen.Render(todos)
}

func BuildHelpMessage(m *RootModel) string {
	return "Help Message Line 1\nHelp Message Line 2"
}

func BuildHomeScreen(m *RootModel) string {
	helpText := BuildHelpMessage(m)

	return lipgloss.JoinVertical(lipgloss.Left, BuildTodosList(m), helpText)
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
