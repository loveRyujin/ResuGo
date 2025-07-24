package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// StartCreateResume starts the interactive resume creation interface
func StartCreateResume() error {
	p := tea.NewProgram(NewModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
