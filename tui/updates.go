package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m RootModel) Init() tea.Cmd {
	return nil
}

type FetchTodosMsg struct{}

type ErrorScreenMsg struct{}

func FetchTodos(m *RootModel) tea.Cmd {
	return func() tea.Msg {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:3000/api/todos", nil)
		if err != nil {
			return ErrorScreenMsg{}
		}
		req.Header.Add("Authorization", "Bearer "+m.Auth.Token)
		res, err := client.Do(req)
		if err != nil {
			return ErrorScreenMsg{}
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return ErrorScreenMsg{}
		}
		fmt.Println(string(body))

		return FetchTodosMsg{}
	}
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

	case FetchTodosMsg:
		m.Screen = "HOME_VIEW"
		return m, nil

	case ErrorScreenMsg:
		m.Screen = "ERROR_VIEW"
		return m, nil

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
				return m, tea.Batch(m.EntryView.spinner.Tick, FetchTodos(&m))
			}
		}
		if m.Screen == ErrorScreen {
			return m, tea.Quit
		}
	}
	return m, nil
}