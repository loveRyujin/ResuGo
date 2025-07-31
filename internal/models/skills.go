package models

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
