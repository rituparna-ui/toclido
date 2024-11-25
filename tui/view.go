package main

func (m RootModel) View() string {
	if m.Screen == EntryScreen {
		return BuildWelcomeScreen(&m)
	}
	if m.Screen == HomeScreen {
		return BuildHomeScreen(&m)
	}
	return "Loading..."
}
