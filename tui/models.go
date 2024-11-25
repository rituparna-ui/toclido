package main

import "github.com/charmbracelet/bubbles/progress"

const (
	EntryScreen Screen = "ENTRY_VIEW"
	HomeScreen  Screen = "HOME_VIEW"
)

type Window struct {
	Width  int
	Height int
}

type Screen string

type EntryView struct {
	showLoader bool
	progress   progress.Model
}

type RootModel struct {
	Token string

	Window Window
	Screen Screen

	EntryView EntryView
}

func NewModel() RootModel {
	return RootModel{
		Screen: EntryScreen,
	}
}
