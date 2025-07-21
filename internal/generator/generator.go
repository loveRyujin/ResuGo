package generator

import (
	"fmt"
	"os"
	"path/filepath"

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
	var content string
	r := g.resume

	// Header
	content += fmt.Sprintf("# %s\n\n", r.PersonalInfo.Name)
	if r.PersonalInfo.Title != "" {
		content += fmt.Sprintf("**%s**\n\n", r.PersonalInfo.Title)
	}

	// Contact Information
	content += "## Contact Information\n\n"
	if r.PersonalInfo.Email != "" {
		content += fmt.Sprintf("- **Email:** %s\n", r.PersonalInfo.Email)
	}
	if r.PersonalInfo.Phone != "" {
		content += fmt.Sprintf("- **Phone:** %s\n", r.PersonalInfo.Phone)
	}
	if r.PersonalInfo.Location != "" {
		content += fmt.Sprintf("- **Location:** %s\n", r.PersonalInfo.Location)
	}
	if r.PersonalInfo.Website != "" {
		content += fmt.Sprintf("- **Website:** %s\n", r.PersonalInfo.Website)
	}
	if r.PersonalInfo.GitHub != "" {
		content += fmt.Sprintf("- **GitHub:** %s\n", r.PersonalInfo.GitHub)
	}
	if r.PersonalInfo.LinkedIn != "" {
		content += fmt.Sprintf("- **LinkedIn:** %s\n", r.PersonalInfo.LinkedIn)
	}
	content += "\n"

	// Summary
	if r.PersonalInfo.Summary != "" {
		content += "## Summary\n\n"
		content += fmt.Sprintf("%s\n\n", r.PersonalInfo.Summary)
	}

	// Education
	if len(r.Education) > 0 {
		content += "## Education\n\n"
		for _, edu := range r.Education {
			content += fmt.Sprintf("### %s\n", edu.Institution)
			content += fmt.Sprintf("**%s in %s**\n", edu.Degree, edu.Major)
			content += fmt.Sprintf("*%s - %s*\n", edu.StartDate.Format("2006"), edu.EndDate.Format("2006"))
			if edu.GPA != "" {
				content += fmt.Sprintf("GPA: %s\n", edu.GPA)
			}
			if edu.Description != "" {
				content += fmt.Sprintf("\n%s\n", edu.Description)
			}
			content += "\n"
		}
	}

	// Experience
	if len(r.Experience) > 0 {
		content += "## Work Experience\n\n"
		for _, exp := range r.Experience {
			content += fmt.Sprintf("### %s\n", exp.Position)
			content += fmt.Sprintf("**%s** - %s\n", exp.Company, exp.Location)
			endDate := "Present"
			if !exp.Current {
				endDate = exp.EndDate.Format("Jan 2006")
			}
			content += fmt.Sprintf("*%s - %s*\n\n", exp.StartDate.Format("Jan 2006"), endDate)

			for _, desc := range exp.Description {
				content += fmt.Sprintf("- %s\n", desc)
			}
			content += "\n"
		}
	}

	// Skills
	if len(r.Skills) > 0 {
		content += "## Skills\n\n"
		for _, skill := range r.Skills {
			content += fmt.Sprintf("### %s", skill.Category)
			if skill.Level != "" {
				content += fmt.Sprintf(" (%s)", skill.Level)
			}
			content += "\n"
			for _, item := range skill.Items {
				content += fmt.Sprintf("- %s\n", item)
			}
			content += "\n"
		}
	}

	// Projects
	if len(r.Projects) > 0 {
		content += "## Projects\n\n"
		for _, project := range r.Projects {
			content += fmt.Sprintf("### %s\n", project.Name)
			content += fmt.Sprintf("%s\n", project.Description)
			content += fmt.Sprintf("*%s - %s*\n", project.StartDate, project.EndDate)

			if len(project.Technologies) > 0 {
				content += "\n**Technologies:** "
				for i, tech := range project.Technologies {
					if i > 0 {
						content += ", "
					}
					content += tech
				}
				content += "\n"
			}

			if project.URL != "" {
				content += fmt.Sprintf("\n**URL:** %s\n", project.URL)
			}

			if project.Repository != "" {
				content += fmt.Sprintf("**Repository:** %s\n", project.Repository)
			}

			if len(project.Highlights) > 0 {
				content += "\n**Highlights:**\n"
				for _, highlight := range project.Highlights {
					content += fmt.Sprintf("- %s\n", highlight)
				}
			}
			content += "\n"
		}
	}

	// Languages
	if len(r.Languages) > 0 {
		content += "## Languages\n\n"
		for _, lang := range r.Languages {
			content += fmt.Sprintf("- **%s:** %s\n", lang.Name, lang.Level)
		}
		content += "\n"
	}

	return content
}
