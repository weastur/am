package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate documentation",
	Long:  `Generate documentation for the am.`,
}

func init() {
	rootCmd.AddCommand(genCmd)
}
