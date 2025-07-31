package models

import "time"

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

// FormatStartDate formats the start date for display
func (e *Experience) FormatStartDate() string {
	return e.StartDate.Format("Jan 2006")
}

// FormatEndDate formats the end date for display
func (e *Experience) FormatEndDate() string {
	if e.Current {
		return "Present"
	}
	return e.EndDate.Format("Jan 2006")
}
