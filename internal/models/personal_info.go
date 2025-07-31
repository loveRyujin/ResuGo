package models

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
