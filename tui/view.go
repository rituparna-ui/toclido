package main

func (m RootModel) View() string {
	if m.Screen == EntryScreen {
		return WelcomeScreen(&m)
	}
	return "Loading..."
}
