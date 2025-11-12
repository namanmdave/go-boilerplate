package cmd

import (
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/server/http"

	"github.com/spf13/cobra"
)

// startHTTPServerCmd represents the start-http-server command
var startHTTPServerCmd = &cobra.Command{
	Use:   "start-http-server",
	Short: "Start the HTTP server",
	Long:  `Start the HTTP server and listen for incoming requests.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting HTTP server...")
		httpServer := http.InitHTTPServer()
		if err := httpServer.Start(); err != nil {
			return fmt.Errorf("failed to start HTTP server: %w", err)
		}
		return nil
	},
}

func init() {
	AddCommand(startHTTPServerCmd)
}
