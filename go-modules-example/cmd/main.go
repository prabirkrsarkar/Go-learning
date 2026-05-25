package cmd

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	temp int
}

func (m model) Init() tea.Cmd {
	return nil
}
