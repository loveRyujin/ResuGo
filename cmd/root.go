package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resumgo",
	Short: "A command-line resume generation tool",
	Long: `ResuGo is an interactive command-line tool for creating and managing personal resumes.
It uses elegant terminal user interface to help you build professional resumes.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add global flags here
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file path")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show verbose output")
}
