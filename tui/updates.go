package main

import tea "github.com/charmbracelet/bubbletea"

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
