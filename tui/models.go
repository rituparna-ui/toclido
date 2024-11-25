package main

import (
	"github.com/charmbracelet/bubbles/spinner"
)

const (
	EntryScreen Screen = "ENTRY_VIEW"
	HomeScreen  Screen = "HOME_VIEW"
	ErrorScreen Screen = "ERROR_VIEW"
)

type Window struct {
	Width  int
	Height int
}

type Screen string

type EntryView struct {
	showLoader bool
	spinner    spinner.Model
}

type Auth struct {
	Token string
	Name  string
	Email string
}

type Todo struct {
	Id          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"createdAt"`
}

type HomeView struct {
	Todos []Todo
}

type RootModel struct {
	Auth Auth

	Window Window
	Screen Screen

	EntryView EntryView
	HomeView  HomeView
}

func NewModel() RootModel {
	return RootModel{
		Screen: EntryScreen,
		HomeView: HomeView{
			Todos: []Todo{},
		},
	}
}
