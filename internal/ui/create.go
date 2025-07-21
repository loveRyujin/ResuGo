package ui

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/loveRyujin/ResuGo/internal/models"
)

// Model represents the application state
type Model struct {
	currentStep int
	resume      models.Resume
	input       string
	cursor      int
	choices     []string
	selected    map[int]struct{}
	quitting    bool
	finished    bool
}

// Define resume creation steps
const (
	StepWelcome = iota
	StepPersonalInfo
	StepEducation
	StepExperience
	StepSkills
	StepProjects
	StepLanguages
	StepFinish
)

func initialModel() Model {
	return Model{
		currentStep: StepWelcome,
		selected:    make(map[int]struct{}),
		choices: []string{
			"Start creating resume",
			"View example",
			"Exit",
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if m.currentStep == StepWelcome {
				if len(m.choices) > 0 && m.cursor == 0 {
					m.currentStep = StepPersonalInfo
				} else if m.cursor == 2 {
					m.quitting = true
					return m, tea.Quit
				}
			} else if m.currentStep == StepFinish {
				m.finished = true
				return m, tea.Quit
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return "Goodbye! 👋\n"
	}

	if m.finished {
		return "🎉 Resume created successfully!\nResume saved to resume.yaml\n"
	}

	var s strings.Builder

	switch m.currentStep {
	case StepWelcome:
		s.WriteString("✨ Welcome to ResuGo Resume Generator ✨\n\n")
		s.WriteString("Please select an option:\n\n")

		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

	case StepPersonalInfo:
		s.WriteString("📝 Personal Information\n\n")
		s.WriteString("This will display personal information input form\n")
		s.WriteString("Press Enter to continue to next step...\n")

	case StepEducation:
		s.WriteString("🎓 Education\n\n")
		s.WriteString("This will display education input form\n")

	case StepExperience:
		s.WriteString("💼 Work Experience\n\n")
		s.WriteString("This will display work experience input form\n")

	case StepSkills:
		s.WriteString("🛠️ Skills\n\n")
		s.WriteString("This will display skills input form\n")

	case StepProjects:
		s.WriteString("🚀 Projects\n\n")
		s.WriteString("This will display projects input form\n")

	case StepLanguages:
		s.WriteString("🌍 Languages\n\n")
		s.WriteString("This will display languages input form\n")

	default:
		s.WriteString("Unknown step\n")
	}

	s.WriteString("\nPress 'q' or Ctrl+C to exit\n")
	return s.String()
}

// StartCreateResume starts the interactive resume creation interface
func StartCreateResume() error {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
