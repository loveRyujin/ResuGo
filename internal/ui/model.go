package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
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

	// Experience/Project management
	managingExperiences bool
	managingProjects    bool
	editingExperience   int // -1 for new, >= 0 for editing existing
	editingProject      int // -1 for new, >= 0 for editing existing
	selectedExperience  int // Currently selected experience in management list
	selectedProject     int // Currently selected project in management list

	// Bubbles components
	welcomeList list.Model
	textInputs  []textinput.Model
	textArea    textarea.Model
	progressBar progress.Model
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

	// Create textarea for multiline inputs
	ta := textarea.New()
	ta.Placeholder = "请输入内容..."
	ta.Focus()
	ta.CharLimit = 500
	ta.SetWidth(60)
	ta.SetHeight(5)

	// Create progress bar
	prog := progress.New(progress.WithDefaultGradient())
	prog.Width = 40

	return Model{
		currentStep: StepWelcome,
		choices: []string{
			"开始创建简历",
			"查看示例",
			"退出",
		},
		customSections: []CustomSection{},
		welcomeList:    welcomeList,
		textInputs:     []textinput.Model{},
		textArea:       ta,
		progressBar:    prog,
	}
}

// createTextInputs creates text input components for the current step fields
func (m *Model) createTextInputs() {
	m.textInputs = make([]textinput.Model, len(m.fields))

	for i, field := range m.fields {
		ti := textinput.New()
		ti.Placeholder = field.Placeholder
		ti.SetValue(field.Value)
		ti.CharLimit = 156
		ti.Width = 50

		if !field.Multiline && !field.IsList {
			// Focus the first single-line input
			if i == 0 {
				ti.Focus()
			}
		} else {
			// For multiline/list fields, create but don't focus
			ti.Blur()
		}

		m.textInputs[i] = ti
	}
}

// Init initializes the model (required by BubbleTea)
func (m Model) Init() tea.Cmd {
	return nil
}

// Update processes messages and updates the model (required by BubbleTea)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Handle window size changes
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.welcomeList.SetWidth(msg.Width)
		m.welcomeList.SetHeight(msg.Height - 4) // Leave space for padding

		// Update textarea width
		m.textArea.SetWidth(msg.Width - 10)

		// Update progress bar width
		m.progressBar.Width = msg.Width - 20 // Leave some padding
		if m.progressBar.Width < 20 {
			m.progressBar.Width = 20 // Minimum width
		}
	}

	// Update welcome list if we're on welcome step
	if m.currentStep == StepWelcome {
		m.welcomeList, cmd = m.welcomeList.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.currentStep != StepWelcome && m.currentStep != StepConfirm && len(m.fields) > 0 {
		// Update text inputs and textarea for form steps
		if !m.editingList {
			// Update all text inputs
			for i := range m.textInputs {
				if i < len(m.fields) && !m.fields[i].Multiline && !m.fields[i].IsList {
					m.textInputs[i], cmd = m.textInputs[i].Update(msg)
					cmds = append(cmds, cmd)
					// Sync textinput value back to field
					m.fields[i].Value = m.textInputs[i].Value()
				}
			}

			// Update textarea for multiline fields
			if m.currentField < len(m.fields) && m.fields[m.currentField].Multiline {
				m.textArea, cmd = m.textArea.Update(msg)
				cmds = append(cmds, cmd)
				// Sync textarea value back to field
				m.fields[m.currentField].Value = m.textArea.Value()
			}
		}
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
			// If we're in edit mode for experience/project, return to management
			if m.editingExperience >= 0 || m.editingProject >= 0 {
				if m.editingExperience >= 0 {
					m.cancelExperienceEdit() // Return to experience management
				} else {
					m.cancelProjectEdit() // Return to project management
				}
				return m, nil
			}
			if m.currentStep > StepWelcome {
				m.currentStep--
				m.setupStep()
			}
			return m, nil

		case "enter":
			return m.handleEnter()

		case "up":
			// Handle navigation in management modes
			if m.managingExperiences {
				if len(m.resume.Experience) > 0 {
					m.selectedExperience = (m.selectedExperience - 1 + len(m.resume.Experience)) % len(m.resume.Experience)
				}
				return m, nil
			}
			if m.managingProjects {
				if len(m.resume.Projects) > 0 {
					m.selectedProject = (m.selectedProject - 1 + len(m.resume.Projects)) % len(m.resume.Projects)
				}
				return m, nil
			}
			// Arrow keys always handle navigation to switch fields
			if m.currentStep != StepWelcome && !m.editingList {
				m.handleListNavigation("up")
				m.handleFormNavigation("up")
			}

		case "down":
			// Handle navigation in management modes
			if m.managingExperiences {
				if len(m.resume.Experience) > 0 {
					m.selectedExperience = (m.selectedExperience + 1) % len(m.resume.Experience)
				}
				return m, nil
			}
			if m.managingProjects {
				if len(m.resume.Projects) > 0 {
					m.selectedProject = (m.selectedProject + 1) % len(m.resume.Projects)
				}
				return m, nil
			}
			// Arrow keys always handle navigation to switch fields
			if m.currentStep != StepWelcome && !m.editingList {
				m.handleListNavigation("down")
				m.handleFormNavigation("down")
			}

		case "tab":
			// Handle tab in management modes
			if m.managingExperiences || m.managingProjects {
				// Continue to next step from management
				m.nextStep()
				return m, nil
			}
			// Always handle tab navigation to switch fields
			if m.currentStep != StepWelcome && !m.editingList {
				m.handleFormNavigation("tab")
			}

		case "shift+tab":
			// Always handle shift+tab navigation to switch fields
			if m.currentStep != StepWelcome && !m.editingList {
				m.handleFormNavigation("shift+tab")
			}

		case "backspace":
			if m.editingList {
				m.handleBackspace()
			}

		case "delete":
			m.handleListManagement("delete")

		case "ctrl+n":
			m.handleListManagement("add")

		case "n", "N":
			// Handle adding new experience/project in management mode
			if m.managingExperiences {
				m.enterExperienceEditMode(-1) // Add new experience
				return m, nil
			}
			if m.managingProjects {
				m.enterProjectEditMode(-1) // Add new project
				return m, nil
			}

		default:
			if m.editingList {
				m.handleTextInput(msg)
			}
		}
	}

	return m, tea.Batch(cmds...)
}

// isCurrentFieldMultiline checks if the current field is multiline
func (m Model) isCurrentFieldMultiline() bool {
	return len(m.fields) > 0 && m.fields[m.currentField].Multiline
}

// calculateProgress returns the current progress percentage (0.0 to 1.0)
func (m Model) calculateProgress() float64 {
	// 总共有 9 个步骤 (Welcome=0, PersonalInfo=1, Summary=2, Education=3, Experience=4, Projects=5, Skills=6, CustomSections=7, Confirm=8, Finish=9)
	totalSteps := float64(9)
	currentStep := float64(m.currentStep)

	// 限制在有效范围内
	if currentStep < 0 {
		currentStep = 0
	}
	if currentStep > totalSteps {
		currentStep = totalSteps
	}

	return currentStep / totalSteps
}

// getStepName returns the Chinese name for current step
func (m Model) getStepName() string {
	stepNames := map[int]string{
		StepWelcome:        "欢迎",
		StepPersonalInfo:   "个人信息",
		StepSummary:        "个人简介",
		StepEducation:      "教育背景",
		StepExperience:     "工作经验",
		StepProjects:       "项目经验",
		StepSkills:         "技能",
		StepCustomSections: "自定义章节",
		StepConfirm:        "确认信息",
		StepFinish:         "完成",
	}

	if name, exists := stepNames[m.currentStep]; exists {
		return name
	}
	return "未知步骤"
}

// blurAllInputs removes focus from all input components
func (m *Model) blurAllInputs() {
	// Blur all textinputs
	for i := range m.textInputs {
		m.textInputs[i].Blur()
	}
	// Blur textarea
	m.textArea.Blur()
}

// focusCurrentField sets focus to the current field's input component
func (m *Model) focusCurrentField() {
	if len(m.fields) == 0 {
		return
	}

	// Blur all inputs first
	m.blurAllInputs()

	currentField := m.fields[m.currentField]

	if currentField.Multiline {
		// Focus textarea for multiline fields
		m.textArea.SetValue(currentField.Value)
		m.textArea.Placeholder = currentField.Placeholder
		m.textArea.Focus()
	} else if !currentField.IsList && m.currentField < len(m.textInputs) {
		// Focus textinput for single-line fields
		m.textInputs[m.currentField].SetValue(currentField.Value)
		m.textInputs[m.currentField].Focus()
	}
}

// syncInputsToFields syncs all input component values back to fields
func (m *Model) syncInputsToFields() {
	if len(m.fields) == 0 {
		return
	}

	for i := range m.fields {
		if m.fields[i].Multiline {
			// Sync textarea value
			m.fields[i].Value = m.textArea.Value()
		} else if !m.fields[i].IsList && i < len(m.textInputs) {
			// Sync textinput value
			m.fields[i].Value = m.textInputs[i].Value()
		}
	}
}
