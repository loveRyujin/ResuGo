package models

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
