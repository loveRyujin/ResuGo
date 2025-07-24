package ui

// Step constants define the steps in the resume creation flow
const (
	StepWelcome = iota
	StepPersonalInfo
	StepSummary
	StepEducation
	StepExperience
	StepProjects
	StepSkills
	StepCustomSections
	StepConfirm
	StepFinish
)

// FormField represents a form input field
type FormField struct {
	Label       string
	Value       string
	Required    bool
	Placeholder string
	Multiline   bool
	IsList      bool // For comma-separated lists
}

// CustomSection represents a user-defined section
type CustomSection struct {
	Title string
	Items []string
}
