package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m RootModel) Init() tea.Cmd {
	return nil
}

type FetchTodosMsg []Todo

type ToggleTodoMsg Todo

type DeleteTodoMsg Todo

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

		var todos struct {
			Todos []Todo `form:"todos"`
		}
		json.Unmarshal(body, &todos)
		m.HomeView.Todos = todos.Todos
		return FetchTodosMsg(m.HomeView.Todos)
	}
}

func ToggleTodo(m *RootModel) tea.Cmd {
	return func() tea.Msg {
		client := &http.Client{}
		url := "http://localhost:3000/api/todos/toggle-status/" + m.HomeView.Todos[m.HomeView.index].Id
		req, err := http.NewRequest("PATCH", url, nil)
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

		var todo struct {
			Todo Todo `form:"todo"`
		}
		json.Unmarshal(body, &todo)
		m.HomeView.Todos[m.HomeView.index] = todo.Todo
		return ToggleTodoMsg{}
	}
}

func DeleteTodo(m *RootModel) tea.Cmd {
	return func() tea.Msg {
		client := &http.Client{}
		url := "http://localhost:3000/api/todos/" + m.HomeView.Todos[m.HomeView.index].Id
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return ErrorScreenMsg{}
		}
		req.Header.Add("Authorization", "Bearer "+m.Auth.Token)
		res, err := client.Do(req)
		if err != nil {
			return ErrorScreenMsg{}
		}
		defer res.Body.Close()

		return DeleteTodoMsg{}
	}
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Width = msg.Width
		m.Window.Height = msg.Height

	case FetchTodosMsg:
		m.Screen = "HOME_VIEW"
		m.HomeView.Todos = msg
		m.EntryView.spinner = spinner.New()
		return m, nil

	case DeleteTodoMsg:
		return m, tea.Cmd(FetchTodos(&m))

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

		if m.Screen == HomeScreen {
			switch msg.String() {
			case "down":
				m.HomeView.index++
				if m.HomeView.index >= len(m.HomeView.Todos) {
					m.HomeView.index = 0
				}
			case "up":
				m.HomeView.index--
				if m.HomeView.index < 0 {
					m.HomeView.index = len(m.HomeView.Todos) - 1
				}
			case "esc":
				return m, tea.Quit
			case " ":
				return m, tea.Cmd(ToggleTodo(&m))
			case "r":
				return m, tea.Cmd(FetchTodos(&m))
			case "d":
				return m, tea.Cmd(DeleteTodo(&m))
			}
		}

		if m.Screen == ErrorScreen {
			return m, tea.Quit
		}
	}
	return m, nil
}
