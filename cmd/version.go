package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ResuGo",
	Long:  "Print the version number of ResuGo resume generator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ResuGo v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
