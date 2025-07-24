package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/loveRyujin/ResuGo/internal/models"
)

// Model represents the application state
type Model struct {
	currentStep  int
	currentField int
	resume       models.Resume
	input        string
	cursor       int
	choices      []string
	quitting     bool
	finished     bool
	error        string

	// Form fields for current step
	fields []FormField

	// Multi-item management
	editingList bool
	listItems   []string
	listIndex   int

	// Custom sections
	customSections []CustomSection
}

// NewModel creates and returns the initial model state
func NewModel() Model {
	return Model{
		currentStep: StepWelcome,
		choices: []string{
			"开始创建简历",
			"查看示例",
			"退出",
		},
		customSections: []CustomSection{},
	}
}

// Init initializes the model (required by BubbleTea)
func (m Model) Init() tea.Cmd {
	return nil
}

// Update processes messages and updates the model (required by BubbleTea)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "esc":
			if m.editingList {
				m.editingList = false
				return m, nil
			}
			if m.currentStep > StepWelcome {
				m.currentStep--
				m.setupStep()
			}
			return m, nil

		case "enter":
			return m.handleEnter()

		case "up", "k":
			m.handleListNavigation("up")
			m.handleFormNavigation("up")

		case "down", "j":
			m.handleListNavigation("down")
			m.handleFormNavigation("down")

		case "tab":
			m.handleFormNavigation("tab")

		case "shift+tab":
			m.handleFormNavigation("shift+tab")

		case "backspace":
			m.handleBackspace()

		case "delete":
			m.handleListManagement("delete")

		case "ctrl+n":
			m.handleListManagement("add")

		default:
			m.handleTextInput(msg)
		}
	}

	return m, nil
}
