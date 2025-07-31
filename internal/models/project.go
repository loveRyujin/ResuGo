package models

import "time"

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

// FormatStartDate formats the start date for display
func (p *Project) FormatStartDate() string {
	return p.StartDate.Format("Jan 2006")
}

// FormatEndDate formats the end date for display
func (p *Project) FormatEndDate() string {
	if p.Current {
		return "Present"
	}
	return p.EndDate.Format("Jan 2006")
}
