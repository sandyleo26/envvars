package cmd

import (
	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/spf13/cobra"
)

var envfileName string
var overwriteEnvfile bool

var envfileCmd = &cobra.Command{
	Use:   "envfile",
	Short: "Generate an env file based on the declaration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := yml.NewDeclaration(declarationFileRootFlag)
		if err != nil {
			return err
		}
		writer := envfile.NewEnvfile(envfileName, overwriteEnvfile)
		return envvars.Envfile(d, writer, tagsRootFlag...)
	},
}

func init() {
	envfileCmd.Flags().StringVar(&envfileName, "env-file", ".env", "env file to be generated")
	envfileCmd.Flags().BoolVar(&overwriteEnvfile, "overwrite", true, "overwrite the env file if it exists")
	rootCmd.AddCommand(envfileCmd)
}
