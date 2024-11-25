package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.EntryView.spinner, cmd = m.EntryView.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

		if m.Screen == EntryScreen {
			switch msg.String() {
			case "enter":
				m.EntryView.showLoader = true
				m.EntryView.spinner.Spinner = spinner.Monkey
				return m, m.EntryView.spinner.Tick
			}
		}
	}
	return m, nil
}
