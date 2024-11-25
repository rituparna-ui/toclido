package main

func (m RootModel) View() string {
	if m.Screen == EntryScreen {
		return BuildWelcomeScreen(&m)
	}
	if m.Screen == HomeScreen {
		return BuildHomeScreen(&m)
	}
	if m.Screen == ErrorScreen {
		return BuildErrorScreen(&m)
	}
	return "Loading..."
}
