package models

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
