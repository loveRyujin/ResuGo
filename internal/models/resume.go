package models

import "time"

// Resume represents a complete resume
type Resume struct {
	PersonalInfo PersonalInfo `yaml:"personal_info"`
	Summary      string       `yaml:"summary,omitempty"`
	Education    []Education  `yaml:"education"`
	Experience   []Experience `yaml:"experience"`
	Projects     []Project    `yaml:"projects"`
	Skills       Skills       `yaml:"skills"`
	Languages    []Language   `yaml:"languages,omitempty"`
	Additional   []Section    `yaml:"additional,omitempty"` // For custom sections
}

// PersonalInfo represents personal basic information
type PersonalInfo struct {
	Name     string `yaml:"name"`
	Title    string `yaml:"title,omitempty"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
	Location string `yaml:"location"`
	Website  string `yaml:"website,omitempty"`
	GitHub   string `yaml:"github,omitempty"`
	LinkedIn string `yaml:"linkedin,omitempty"`
}

// Education represents educational background
type Education struct {
	Institution     string    `yaml:"institution"`
	Degree          string    `yaml:"degree"`
	Major           string    `yaml:"major,omitempty"`
	StartDate       time.Time `yaml:"start_date"`
	EndDate         time.Time `yaml:"end_date"`
	Current         bool      `yaml:"current,omitempty"`
	Location        string    `yaml:"location"`
	GPA             string    `yaml:"gpa,omitempty"`
	RelevantCourses []string  `yaml:"relevant_courses,omitempty"`
	HonorsAwards    []string  `yaml:"honors_awards,omitempty"`
	Description     string    `yaml:"description,omitempty"`
}

// Experience represents work experience
type Experience struct {
	Company          string    `yaml:"company"`
	Position         string    `yaml:"position"`
	Location         string    `yaml:"location"`
	StartDate        time.Time `yaml:"start_date"`
	EndDate          time.Time `yaml:"end_date"`
	Current          bool      `yaml:"current"`
	Responsibilities []string  `yaml:"responsibilities"`
	Achievements     []string  `yaml:"achievements,omitempty"`
}

// Project represents project experience
type Project struct {
	Name         string    `yaml:"name"`
	Description  string    `yaml:"description"`
	StartDate    time.Time `yaml:"start_date"`
	EndDate      time.Time `yaml:"end_date"`
	Current      bool      `yaml:"current,omitempty"`
	Location     string    `yaml:"location,omitempty"`
	Technologies []string  `yaml:"technologies,omitempty"`
	URL          string    `yaml:"url,omitempty"`
	Repository   string    `yaml:"repository,omitempty"`
	Details      []string  `yaml:"details"`
}

// Skills represents all skills grouped by categories
type Skills struct {
	Languages  []string        `yaml:"languages,omitempty"`
	Frameworks []string        `yaml:"frameworks,omitempty"`
	Databases  []string        `yaml:"databases,omitempty"`
	Tools      []string        `yaml:"tools,omitempty"`
	Other      []string        `yaml:"other,omitempty"`
	Custom     []SkillCategory `yaml:"custom,omitempty"` // For custom skill categories
}

// SkillCategory represents a custom skill category
type SkillCategory struct {
	Name  string   `yaml:"name"`
	Items []string `yaml:"items"`
}

// Language represents language proficiency
type Language struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"` // native, fluent, conversational, basic
}

// Section represents additional custom sections
type Section struct {
	Title string   `yaml:"title"`
	Items []string `yaml:"items"`
}

// FormatDate formats a date for display
func (e *Experience) FormatStartDate() string {
	return e.StartDate.Format("Jan 2006")
}

func (e *Experience) FormatEndDate() string {
	if e.Current {
		return "Present"
	}
	return e.EndDate.Format("Jan 2006")
}

func (edu *Education) FormatStartDate() string {
	return edu.StartDate.Format("2006")
}

func (edu *Education) FormatEndDate() string {
	if edu.Current {
		return "Present"
	}
	return edu.EndDate.Format("2006")
}

func (p *Project) FormatStartDate() string {
	return p.StartDate.Format("Jan 2006")
}

func (p *Project) FormatEndDate() string {
	if p.Current {
		return "Present"
	}
	return p.EndDate.Format("Jan 2006")
}
