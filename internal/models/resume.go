package models

import "time"

// Resume represents a complete resume
type Resume struct {
	PersonalInfo PersonalInfo `yaml:"personal_info"`
	Education    []Education  `yaml:"education"`
	Experience   []Experience `yaml:"experience"`
	Skills       []Skill      `yaml:"skills"`
	Projects     []Project    `yaml:"projects"`
	Languages    []Language   `yaml:"languages"`
}

// PersonalInfo represents personal basic information
type PersonalInfo struct {
	Name     string `yaml:"name"`
	Title    string `yaml:"title"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
	Location string `yaml:"location"`
	Website  string `yaml:"website"`
	GitHub   string `yaml:"github"`
	LinkedIn string `yaml:"linkedin"`
	Summary  string `yaml:"summary"`
}

// Education represents educational background
type Education struct {
	Institution string    `yaml:"institution"`
	Degree      string    `yaml:"degree"`
	Major       string    `yaml:"major"`
	StartDate   time.Time `yaml:"start_date"`
	EndDate     time.Time `yaml:"end_date"`
	GPA         string    `yaml:"gpa,omitempty"`
	Description string    `yaml:"description,omitempty"`
}

// Experience represents work experience
type Experience struct {
	Company     string    `yaml:"company"`
	Position    string    `yaml:"position"`
	Location    string    `yaml:"location"`
	StartDate   time.Time `yaml:"start_date"`
	EndDate     time.Time `yaml:"end_date"`
	Current     bool      `yaml:"current"`
	Description []string  `yaml:"description"`
}

// Skill represents skills
type Skill struct {
	Category string   `yaml:"category"`
	Items    []string `yaml:"items"`
	Level    string   `yaml:"level,omitempty"` // beginner, intermediate, advanced, expert
}

// Project represents project experience
type Project struct {
	Name         string   `yaml:"name"`
	Description  string   `yaml:"description"`
	StartDate    string   `yaml:"start_date"`
	EndDate      string   `yaml:"end_date"`
	Technologies []string `yaml:"technologies"`
	URL          string   `yaml:"url,omitempty"`
	Repository   string   `yaml:"repository,omitempty"`
	Highlights   []string `yaml:"highlights"`
}

// Language represents language proficiency
type Language struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"` // native, fluent, good, basic
}
