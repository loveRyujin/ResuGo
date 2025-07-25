package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/loveRyujin/ResuGo/internal/models"
)

// listItem implements list.Item interface for the welcome list
type listItem struct {
	title string
	desc  string
}

func (i listItem) FilterValue() string { return i.title }
func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return i.desc }

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

	// Bubbles components
	welcomeList list.Model
}

// NewModel creates and returns the initial model state
func NewModel() Model {
	// Create list items for welcome screen
	items := []list.Item{
		listItem{title: "开始创建简历", desc: "创建一份新的简历"},
		listItem{title: "查看示例", desc: "查看简历模板示例"},
		listItem{title: "退出", desc: "退出程序"},
	}

	// Create and configure the list
	welcomeList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	welcomeList.Title = "✨ 欢迎使用 ResuGo 简历生成工具 ✨"
	welcomeList.SetShowStatusBar(false)
	welcomeList.SetFilteringEnabled(false)
	welcomeList.Styles.Title = welcomeList.Styles.Title.Bold(true)

	return Model{
		currentStep: StepWelcome,
		choices: []string{
			"开始创建简历",
			"查看示例",
			"退出",
		},
		customSections: []CustomSection{},
		welcomeList:    welcomeList,
	}
}

// Init initializes the model (required by BubbleTea)
func (m Model) Init() tea.Cmd {
	return nil
}

// Update processes messages and updates the model (required by BubbleTea)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Handle window size changes
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.welcomeList.SetWidth(msg.Width)
		m.welcomeList.SetHeight(msg.Height - 4) // Leave space for padding
	}

	// Update welcome list if we're on welcome step
	if m.currentStep == StepWelcome {
		m.welcomeList, cmd = m.welcomeList.Update(msg)
	}

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
			if m.currentStep != StepWelcome {
				m.handleListNavigation("up")
				m.handleFormNavigation("up")
			}

		case "down", "j":
			if m.currentStep != StepWelcome {
				m.handleListNavigation("down")
				m.handleFormNavigation("down")
			}

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

	return m, cmd
}
