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

	if len(m.HomeView.Todos) == 0 {
		todos = "No Todos\nPress [a] to add a new todo"
		return fullScreen.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render(todos)
	}

	for i, todo := range m.HomeView.Todos {
		x := " "
		if todo.Completed {
			x = "x"
		}
		titleForeground := lipgloss.Color("#FF0000")
		if todo.Completed {
			titleForeground = lipgloss.Color("#00FF00")
		}
		todoItem := fmt.Sprintf("[%s] %s\n    %s",
			x,
			lipgloss.
				NewStyle().
				Strikethrough(todo.Completed).
				Foreground(titleForeground).
				Render(todo.Title), todo.Description)

		isActive := m.HomeView.index == i

		leftPad := 1
		if isActive {
			leftPad = 0
		}

		todos += lipgloss.
			NewStyle().
			Width(m.Window.Width-4).
			Border(lipgloss.RoundedBorder(), isActive).
			BorderForeground(lipgloss.Color("#808000")).
			Padding(leftPad, 0, leftPad, leftPad).
			Render(todoItem)
		todos += "\n"
	}

	return fullScreen.Render(todos)
}

func BuildHelpMessage(m *RootModel) string {
	line1 := "[up/down] to navigate          [d] to delete a todo  [r]   to refresh"
	line2 := "[space]   to toggle completed  [a] to add a todo     [esc] to exit"

	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
	line1 = style.Render(line1)
	line2 = style.Render(line2)

	return line1 + "\n" + line2
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
