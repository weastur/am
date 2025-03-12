package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const schemaFilePerms = 0o600

var outputFilename string

//go:embed schema.json
var schema string

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Generate JSON schema",
	Long: `This command generates a JSON schema for the am CLI.
By default, it creates the schema.json file in the current directory.`,
	Run: func(_ *cobra.Command, _ []string) {
		err := os.WriteFile(outputFilename, []byte(schema), schemaFilePerms)
		if err != nil {
			cobra.CheckErr(err)
		}
		fmt.Println("Schema written to", outputFilename)
	},
}

func init() {
	configCmd.AddCommand(schemaCmd)

	schemaCmd.Flags().StringVar(&outputFilename, "out", "config.schema.json", "Filename to write the JSON schema to")
	schemaCmd.MarkFlagFilename("out", "json")
}
