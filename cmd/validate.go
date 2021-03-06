package cmd

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check if the declaration file contains any error",
	Long:  "The flag tags has no effect with this command",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := yml.NewDeclarationYML(declarationFileRootFlag)
		return envvars.Validate(reader)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
