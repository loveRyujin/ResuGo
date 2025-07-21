package cmd

import (
	"fmt"
	"os"

	"github.com/loveRyujin/ResuGo/internal/generator"
	"github.com/loveRyujin/ResuGo/internal/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	outputFormat string
	outputPath   string
)

var generateCmd = &cobra.Command{
	Use:   "generate [input-file]",
	Short: "Generate resume from YAML file",
	Long:  "Generate resume in different formats (markdown, pdf) from a YAML input file",
	Args:  cobra.ExactArgs(1),
	RunE:  generateResume,
}

func generateResume(cmd *cobra.Command, args []string) error {
	inputFile := args[0]

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file %s does not exist", inputFile)
	}

	// Read YAML file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse YAML
	var resume models.Resume
	if err := yaml.Unmarshal(data, &resume); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Create generator
	gen := generator.NewGenerator(&resume)

	// Generate output
	switch outputFormat {
	case "yaml":
		if outputPath == "" {
			outputPath = "resume.yaml"
		}
		if err := gen.GenerateYAML(outputPath); err != nil {
			return fmt.Errorf("failed to generate YAML: %w", err)
		}
	case "markdown", "md":
		if outputPath == "" {
			outputPath = "resume.md"
		}
		if err := gen.GenerateMarkdown(outputPath); err != nil {
			return fmt.Errorf("failed to generate Markdown: %w", err)
		}
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}

	fmt.Printf("Resume generated successfully: %s\n", outputPath)
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)
	
	generateCmd.Flags().StringVarP(&outputFormat, "format", "f", "markdown", "Output format (yaml, markdown)")
	generateCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path")
}
