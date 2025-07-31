package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/loveRyujin/ResuGo/internal/generator"
	"github.com/loveRyujin/ResuGo/internal/models"
)

// saveCurrentStep saves the current step's data to the resume model
func (m *Model) saveCurrentStep() {
	switch m.currentStep {
	case StepPersonalInfo:
		m.savePersonalInfo()
	case StepSummary:
		m.saveSummary()
	case StepEducation:
		m.saveEducation()
	case StepExperience:
		m.saveExperience()
	case StepProjects:
		m.saveProjects()
	case StepSkills:
		m.saveSkills()
	case StepCustomSections:
		m.saveCustomSections()
	}
}

// savePersonalInfo saves personal information data
func (m *Model) savePersonalInfo() {
	m.resume.PersonalInfo = models.PersonalInfo{
		Name:     strings.TrimSpace(m.fields[0].Value),
		Email:    strings.TrimSpace(m.fields[1].Value),
		Phone:    strings.TrimSpace(m.fields[2].Value),
		Location: strings.TrimSpace(m.fields[3].Value),
		Website:  strings.TrimSpace(m.fields[4].Value),
	}
}

// saveSummary saves summary data
func (m *Model) saveSummary() {
	m.resume.Summary = strings.TrimSpace(m.fields[0].Value)
}

// saveEducation saves education data
func (m *Model) saveEducation() {
	var startYear, endYear int
	var err error

	startYear, err = strconv.Atoi(m.fields[4].Value)
	if err != nil {
		startYear = 2020 // default
	}

	if m.fields[5].Value == "current" {
		endYear = time.Now().Year()
	} else {
		endYear, err = strconv.Atoi(m.fields[5].Value)
		if err != nil {
			endYear = startYear + 4 // default
		}
	}

	edu := models.Education{
		Institution: strings.TrimSpace(m.fields[0].Value),
		Degree:      strings.TrimSpace(m.fields[1].Value),
		Major:       strings.TrimSpace(m.fields[2].Value),
		Location:    strings.TrimSpace(m.fields[3].Value),
		StartDate:   time.Date(startYear, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:     time.Date(endYear, 5, 1, 0, 0, 0, 0, time.UTC),
		Current:     m.fields[5].Value == "current",
	}

	// Update existing education or add new one
	if len(m.resume.Education) > 0 {
		// Update the first education entry
		m.resume.Education[0] = edu
	} else {
		// Add new education entry
		m.resume.Education = append(m.resume.Education, edu)
	}
}

// saveExperience saves work experience data and returns to management
func (m *Model) saveExperience() {
	startDate, _ := time.Parse("2006-01", m.fields[3].Value)
	var endDate time.Time
	var current bool

	if m.fields[4].Value == "current" {
		current = true
		endDate = time.Now()
	} else {
		endDate, _ = time.Parse("2006-01", m.fields[4].Value)
	}

	// Split responsibilities by newlines and filter empty ones
	responsibilities := []string{}
	for _, line := range strings.Split(strings.TrimSpace(m.fields[5].Value), "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			responsibilities = append(responsibilities, trimmed)
		}
	}

	exp := models.Experience{
		Company:          strings.TrimSpace(m.fields[0].Value),
		Position:         strings.TrimSpace(m.fields[1].Value),
		Location:         strings.TrimSpace(m.fields[2].Value),
		StartDate:        startDate,
		EndDate:          endDate,
		Current:          current,
		Responsibilities: responsibilities,
	}

	// Update existing experience or add new one based on editingExperience index
	if m.editingExperience >= 0 && m.editingExperience < len(m.resume.Experience) {
		// Update existing experience
		m.resume.Experience[m.editingExperience] = exp
	} else {
		// Add new experience entry
		m.resume.Experience = append(m.resume.Experience, exp)
	}

	// Return to experience management using dedicated method
	wasEditing := m.editingExperience >= 0
	editIndex := m.editingExperience
	m.returnToExperienceManagement(wasEditing, editIndex)
}

// returnToExperienceManagement returns to experience management mode with proper state
func (m *Model) returnToExperienceManagement(wasEditing bool, editIndex int) {
	// Set management mode states
	m.managingExperiences = true
	m.editingExperience = -1
	m.fields = nil

	// Set proper selection
	if wasEditing {
		m.selectedExperience = editIndex
	} else {
		// Select the newly added item (last in list)
		m.selectedExperience = len(m.resume.Experience) - 1
	}

	// Ensure selection is valid
	if m.selectedExperience >= len(m.resume.Experience) {
		m.selectedExperience = len(m.resume.Experience) - 1
	}
	if m.selectedExperience < 0 {
		m.selectedExperience = 0
	}

	// Clear any form-related states
	m.currentField = 0
	m.error = ""
	m.editingList = false
}

// cancelExperienceEdit returns to experience management mode when canceling edit
func (m *Model) cancelExperienceEdit() {
	// Set management mode states
	m.managingExperiences = true
	m.editingExperience = -1
	m.fields = nil

	// Keep current selection or reset to 0 if invalid
	if m.selectedExperience >= len(m.resume.Experience) {
		m.selectedExperience = 0
	}
	if m.selectedExperience < 0 && len(m.resume.Experience) > 0 {
		m.selectedExperience = 0
	}

	// Clear any form-related states
	m.currentField = 0
	m.error = ""
	m.editingList = false
}

// saveProjects saves project data and returns to management
func (m *Model) saveProjects() {
	startDate, _ := time.Parse("2006-01", m.fields[3].Value)
	var endDate time.Time
	var current bool

	if m.fields[4].Value == "current" {
		current = true
		endDate = time.Now()
	} else {
		endDate, _ = time.Parse("2006-01", m.fields[4].Value)
	}

	// Split details by newlines and filter empty ones
	details := []string{}
	for _, line := range strings.Split(strings.TrimSpace(m.fields[5].Value), "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			details = append(details, trimmed)
		}
	}

	project := models.Project{
		Name:        strings.TrimSpace(m.fields[0].Value),
		Description: strings.TrimSpace(m.fields[1].Value),
		Location:    strings.TrimSpace(m.fields[2].Value),
		StartDate:   startDate,
		EndDate:     endDate,
		Current:     current,
		Details:     details,
	}

	// Update existing project or add new one based on editingProject index
	if m.editingProject >= 0 && m.editingProject < len(m.resume.Projects) {
		// Update existing project
		m.resume.Projects[m.editingProject] = project
	} else {
		// Add new project entry
		m.resume.Projects = append(m.resume.Projects, project)
	}

	// Return to project management using dedicated method
	wasEditing := m.editingProject >= 0
	editIndex := m.editingProject
	m.returnToProjectManagement(wasEditing, editIndex)
}

// returnToProjectManagement returns to project management mode with proper state
func (m *Model) returnToProjectManagement(wasEditing bool, editIndex int) {
	// Set management mode states
	m.managingProjects = true
	m.editingProject = -1
	m.fields = nil

	// Set proper selection
	if wasEditing {
		m.selectedProject = editIndex
	} else {
		// Select the newly added item (last in list)
		m.selectedProject = len(m.resume.Projects) - 1
	}

	// Ensure selection is valid
	if m.selectedProject >= len(m.resume.Projects) {
		m.selectedProject = len(m.resume.Projects) - 1
	}
	if m.selectedProject < 0 {
		m.selectedProject = 0
	}

	// Clear any form-related states
	m.currentField = 0
	m.error = ""
	m.editingList = false
}

// cancelProjectEdit returns to project management mode when canceling edit
func (m *Model) cancelProjectEdit() {
	// Set management mode states
	m.managingProjects = true
	m.editingProject = -1
	m.fields = nil

	// Keep current selection or reset to 0 if invalid
	if m.selectedProject >= len(m.resume.Projects) {
		m.selectedProject = 0
	}
	if m.selectedProject < 0 && len(m.resume.Projects) > 0 {
		m.selectedProject = 0
	}

	// Clear any form-related states
	m.currentField = 0
	m.error = ""
	m.editingList = false
}

// saveSkills saves skills data
func (m *Model) saveSkills() {
	m.resume.Skills = models.Skills{
		Languages:  parseSkillList(m.fields[0].Value),
		Frameworks: parseSkillList(m.fields[1].Value),
		Databases:  parseSkillList(m.fields[2].Value),
		Other:      parseSkillList(m.fields[3].Value),
	}
}

// saveCustomSections saves custom sections data
func (m *Model) saveCustomSections() {
	title := strings.TrimSpace(m.fields[0].Value)
	if title != "" {
		items := parseSkillList(m.fields[1].Value)
		if len(items) > 0 {
			section := models.Section{
				Title: title,
				Items: items,
			}
			m.resume.Additional = append(m.resume.Additional, section)
		}
	}
}

// saveResume saves the complete resume to files
func (m *Model) saveResume() error {
	gen := generator.NewGenerator(&m.resume)

	// Save YAML
	if err := gen.GenerateYAML("my_resume.yaml"); err != nil {
		return fmt.Errorf("保存YAML失败: %w", err)
	}

	// Save Markdown
	if err := gen.GenerateMarkdown("my_resume.md"); err != nil {
		return fmt.Errorf("保存Markdown失败: %w", err)
	}

	return nil
}
