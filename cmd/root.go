package cmd

import (
	"go-boilerplate/config"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boilerplate",
	Short: "Go Boilerplate CLI",
	Long:  `A Go boilerplate application with multiple commands for managing the server and database migrations.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() error {
	return rootCmd.Execute()
}

// AddCommand adds a subcommand to the root command
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
