package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage configuration for the am.`,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
