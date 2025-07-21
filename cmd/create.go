package cmd

import (
	"github.com/loveRyujin/ResuGo/internal/ui"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new resume",
	Long:  "Create a new resume using an interactive interface",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.StartCreateResume()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
