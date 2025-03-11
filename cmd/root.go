package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "am",
	Short: "Asset Manager",
	Long: `Asset Manager is a package manager for software distributed as tarballs or binaries.
It is designed to be simple and easy to use both on local or remotes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.am.yaml)")
}

func initConfig() {
	// var cfg Config = config.Get()

	// cobra.CheckErr(cfg.Init(cfgFile))
}
