package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/loveRyujin/ResuGo/internal/models"
	"gopkg.in/yaml.v3"
)

// Generator handles resume generation in different formats
type Generator struct {
	resume *models.Resume
}

// NewGenerator creates a new generator instance
func NewGenerator(resume *models.Resume) *Generator {
	return &Generator{
		resume: resume,
	}
}

// GenerateYAML generates resume in YAML format
func (g *Generator) GenerateYAML(outputPath string) error {
	data, err := yaml.Marshal(g.resume)
	if err != nil {
		return fmt.Errorf("failed to marshal resume to YAML: %w", err)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write YAML file: %w", err)
	}

	return nil
}

// GenerateMarkdown generates resume in Markdown format
func (g *Generator) GenerateMarkdown(outputPath string) error {
	content := g.buildMarkdownContent()

	// Create directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write Markdown file: %w", err)
	}

	return nil
}

func (g *Generator) buildMarkdownContent() string {
	var content strings.Builder
	r := g.resume

	// Header - Centered name
	content.WriteString(fmt.Sprintf("<div align=\"center\">\n\n# %s\n\n", r.PersonalInfo.Name))

	// Contact information in one line
	var contactParts []string
	if r.PersonalInfo.Phone != "" {
		contactParts = append(contactParts, r.PersonalInfo.Phone)
	}
	if r.PersonalInfo.Email != "" {
		contactParts = append(contactParts, r.PersonalInfo.Email)
	}
	if r.PersonalInfo.Location != "" {
		contactParts = append(contactParts, r.PersonalInfo.Location)
	}
	if r.PersonalInfo.Website != "" {
		contactParts = append(contactParts, r.PersonalInfo.Website)
	}

	if len(contactParts) > 0 {
		content.WriteString(strings.Join(contactParts, " | "))
		content.WriteString("\n\n</div>\n\n")
	} else {
		content.WriteString("</div>\n\n")
	}

	// Summary section
	if r.Summary != "" {
		content.WriteString("## Summary\n\n")
		content.WriteString(fmt.Sprintf("%s\n\n", r.Summary))
		content.WriteString("---\n\n")
	}

	// Education section
	if len(r.Education) > 0 {
		content.WriteString("## Education\n\n")
		for _, edu := range r.Education {
			// Degree line
			content.WriteString(fmt.Sprintf("**%s**", edu.Degree))
			if edu.Major != "" {
				content.WriteString(fmt.Sprintf(" in %s", edu.Major))
			}

			// Institution and dates on right
			content.WriteString(fmt.Sprintf("%s%s - %s\n",
				strings.Repeat(" ", 50), edu.FormatStartDate(), edu.FormatEndDate()))

			// Institution name and location
			content.WriteString(fmt.Sprintf("%s", edu.Institution))
			if edu.Location != "" {
				content.WriteString(fmt.Sprintf("%s%s\n",
					strings.Repeat(" ", 40), edu.Location))
			} else {
				content.WriteString("\n")
			}

			// Additional details
			if len(edu.RelevantCourses) > 0 {
				content.WriteString("• **Relevant Courses:** ")
				content.WriteString(strings.Join(edu.RelevantCourses, ", "))
				content.WriteString("\n")
			}

			if len(edu.HonorsAwards) > 0 {
				content.WriteString("• **Honors & Awards:** ")
				content.WriteString(strings.Join(edu.HonorsAwards, ", "))
				content.WriteString("\n")
			}

			content.WriteString("\n")
		}
		content.WriteString("---\n\n")
	}

	// Experience section
	if len(r.Experience) > 0 {
		content.WriteString("## Experience\n\n")
		for _, exp := range r.Experience {
			// Position and dates
			content.WriteString(fmt.Sprintf("**%s**", exp.Position))
			content.WriteString(fmt.Sprintf("%s%s - %s\n",
				strings.Repeat(" ", 50), exp.FormatStartDate(), exp.FormatEndDate()))

			// Company and location
			content.WriteString(fmt.Sprintf("%s", exp.Company))
			if exp.Location != "" {
				content.WriteString(fmt.Sprintf("%s%s\n",
					strings.Repeat(" ", 40), exp.Location))
			} else {
				content.WriteString("\n")
			}

			// Responsibilities/Description
			for _, resp := range exp.Responsibilities {
				content.WriteString(fmt.Sprintf("• %s\n", resp))
			}

			// Achievements
			for _, achievement := range exp.Achievements {
				content.WriteString(fmt.Sprintf("• **Achievement:** %s\n", achievement))
			}

			content.WriteString("\n")
		}
		content.WriteString("---\n\n")
	}

	// Projects section
	if len(r.Projects) > 0 {
		content.WriteString("## Projects\n\n")
		for _, project := range r.Projects {
			// Project name and dates
			content.WriteString(fmt.Sprintf("**%s**", project.Name))
			content.WriteString(fmt.Sprintf("%s%s - %s\n",
				strings.Repeat(" ", 50), project.FormatStartDate(), project.FormatEndDate()))

			// Description
			content.WriteString(fmt.Sprintf("%s", project.Description))
			if project.Location != "" {
				content.WriteString(fmt.Sprintf("%s%s\n",
					strings.Repeat(" ", 40), project.Location))
			} else {
				content.WriteString("\n")
			}

			// Details
			for _, detail := range project.Details {
				content.WriteString(fmt.Sprintf("• %s\n", detail))
			}

			content.WriteString("\n")
		}
		content.WriteString("---\n\n")
	}

	// Skills section
	content.WriteString("## Skills\n\n")

	if len(r.Skills.Languages) > 0 {
		content.WriteString(fmt.Sprintf("• **Languages:** %s\n",
			strings.Join(r.Skills.Languages, ", ")))
	}

	if len(r.Skills.Frameworks) > 0 {
		content.WriteString(fmt.Sprintf("• **Frameworks:** %s\n",
			strings.Join(r.Skills.Frameworks, ", ")))
	}

	if len(r.Skills.Databases) > 0 {
		content.WriteString(fmt.Sprintf("• **Databases:** %s\n",
			strings.Join(r.Skills.Databases, ", ")))
	}

	if len(r.Skills.Tools) > 0 {
		content.WriteString(fmt.Sprintf("• **Tools:** %s\n",
			strings.Join(r.Skills.Tools, ", ")))
	}

	if len(r.Skills.Other) > 0 {
		content.WriteString(fmt.Sprintf("• **Other:** %s\n",
			strings.Join(r.Skills.Other, ", ")))
	}

	// Custom skill categories
	for _, customSkill := range r.Skills.Custom {
		content.WriteString(fmt.Sprintf("• **%s:** %s\n",
			customSkill.Name, strings.Join(customSkill.Items, ", ")))
	}

	// Languages section (if any)
	if len(r.Languages) > 0 {
		content.WriteString("\n---\n\n## Languages\n\n")
		for _, lang := range r.Languages {
			content.WriteString(fmt.Sprintf("• **%s:** %s\n", lang.Name, lang.Level))
		}
	}

	// Additional sections
	for _, section := range r.Additional {
		content.WriteString(fmt.Sprintf("\n---\n\n## %s\n\n", section.Title))
		for _, item := range section.Items {
			content.WriteString(fmt.Sprintf("• %s\n", item))
		}
	}

	return content.String()
}
