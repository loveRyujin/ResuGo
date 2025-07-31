package models

import "time"

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

// FormatStartDate formats the start date for display
func (edu *Education) FormatStartDate() string {
	return edu.StartDate.Format("2006")
}

// FormatEndDate formats the end date for display
func (edu *Education) FormatEndDate() string {
	if edu.Current {
		return "Present"
	}
	return edu.EndDate.Format("2006")
}
