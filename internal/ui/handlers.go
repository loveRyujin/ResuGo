package ui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
)

// handleEnter processes the Enter key action
func (m *Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.currentStep {
	case StepWelcome:
		if selectedItem := m.welcomeList.SelectedItem(); selectedItem != nil {
			if item, ok := selectedItem.(listItem); ok {
				switch item.title {
				case "开始创建简历":
					m.currentStep = StepPersonalInfo
					m.setupStep()
				case "查看示例":
					// TODO: Implement example view
					return *m, nil
				case "退出":
					m.quitting = true
					return *m, tea.Quit
				}
			}
		}

	case StepConfirm:
		// Save resume
		if err := m.saveResume(); err != nil {
			m.error = fmt.Sprintf("保存失败: %v", err)
			return *m, nil
		}
		m.currentStep = StepFinish
		m.finished = true
		return *m, nil

	default:
		if m.editingList {
			// Exit list editing mode
			m.exitListEditingMode()
			return *m, nil
		}

		// Check if current field is a list field
		if m.enterListEditingMode() {
			return *m, nil
		}

		// Sync all input component values to fields before validation
		m.syncInputsToFields()

		// Validate and save current step data
		if m.validateCurrentStep() {
			m.saveCurrentStep()
			m.nextStep()
		}
	}

	return *m, nil
}

// handleTextInput processes text input for form fields and lists
func (m *Model) handleTextInput(msg tea.KeyMsg) {
	if m.editingList {
		if m.listIndex < len(m.listItems) {
			m.listItems[m.listIndex] += msg.String()
		}
	} else if len(m.fields) > 0 && m.currentStep != StepWelcome {
		// Filter out control characters and ensure valid input
		if utf8.ValidString(msg.String()) && len(msg.String()) == utf8.RuneCountInString(msg.String()) {
			m.fields[m.currentField].Value += msg.String()
		}
	}
}

// handleBackspace processes backspace key for proper Unicode character deletion
func (m *Model) handleBackspace() {
	if m.editingList {
		// Handle list item editing
		if m.listIndex < len(m.listItems) && len(m.listItems[m.listIndex]) > 0 {
			// Use rune-based deletion for proper Unicode support
			runes := []rune(m.listItems[m.listIndex])
			if len(runes) > 0 {
				m.listItems[m.listIndex] = string(runes[:len(runes)-1])
			}
		}
	} else if len(m.fields) > 0 && len(m.fields[m.currentField].Value) > 0 {
		// Use rune-based deletion for proper Unicode support
		runes := []rune(m.fields[m.currentField].Value)
		if len(runes) > 0 {
			m.fields[m.currentField].Value = string(runes[:len(runes)-1])
		}
	}
}

// handleListNavigation processes up/down navigation in lists
func (m *Model) handleListNavigation(direction string) {
	if !m.editingList {
		return
	}

	switch direction {
	case "up":
		if m.listIndex > 0 {
			m.listIndex--
		}
	case "down":
		if m.listIndex < len(m.listItems) {
			m.listIndex++
		}
	}
}

// handleFormNavigation processes form field navigation
func (m *Model) handleFormNavigation(direction string) {
	if m.editingList {
		return
	}

	oldField := m.currentField

	switch direction {
	case "up":
		if m.currentStep == StepWelcome {
			if m.cursor > 0 {
				m.cursor--
			}
		} else if len(m.fields) > 0 {
			if m.currentField > 0 {
				m.currentField--
			}
		}
	case "down":
		if m.currentStep == StepWelcome {
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		} else if len(m.fields) > 0 {
			if m.currentField < len(m.fields)-1 {
				m.currentField++
			}
		}
	case "tab":
		if !m.editingList && len(m.fields) > 0 && m.currentField < len(m.fields)-1 {
			m.currentField++
		}
	case "shift+tab":
		if !m.editingList && len(m.fields) > 0 && m.currentField > 0 {
			m.currentField--
		}
	}

	// Focus the new current field if it changed
	if oldField != m.currentField && m.currentStep != StepWelcome {
		m.focusCurrentField()
	}
}

// handleListManagement processes list editing operations
func (m *Model) handleListManagement(action string) {
	if !m.editingList {
		return
	}

	switch action {
	case "delete":
		if m.listIndex < len(m.listItems) {
			// Delete current list item
			m.listItems = append(m.listItems[:m.listIndex], m.listItems[m.listIndex+1:]...)
			if m.listIndex >= len(m.listItems) && m.listIndex > 0 {
				m.listIndex--
			}
		}
	case "add":
		// Add new list item
		m.listItems = append(m.listItems, "")
		m.listIndex = len(m.listItems) - 1
	}
}

// enterListEditingMode enters the list editing mode for a field
func (m *Model) enterListEditingMode() bool {
	if len(m.fields) == 0 || !m.fields[m.currentField].IsList {
		return false
	}

	m.editingList = true
	m.listItems = parseSkillList(m.fields[m.currentField].Value)

	// If no items exist, create a default empty item
	if len(m.listItems) == 0 {
		m.listItems = []string{""}
	}

	m.listIndex = 0
	return true
}

// exitListEditingMode exits the list editing mode and saves changes
func (m *Model) exitListEditingMode() {
	if !m.editingList {
		return
	}

	m.editingList = false
	// Save list back to field
	if len(m.fields) > 0 {
		// Filter out empty items
		var validItems []string
		for _, item := range m.listItems {
			if trimmed := strings.TrimSpace(item); trimmed != "" {
				validItems = append(validItems, trimmed)
			}
		}
		m.fields[m.currentField].Value = strings.Join(validItems, ", ")
	}
}

// parseSkillList parses a comma-separated string into a slice
func parseSkillList(input string) []string {
	if strings.TrimSpace(input) == "" {
		return nil
	}

	skills := strings.Split(input, ",")
	var result []string
	for _, skill := range skills {
		trimmed := strings.TrimSpace(skill)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
