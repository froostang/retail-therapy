package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/froostang/retail-therapy/api/http"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var template embed.FS

var rootCmd = &cobra.Command{
	Use:   "my-service",
	Short: "A Cobra-based service daemon with Zap logging options",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		logger.Info("Service starting")

		http.NewServer(logger, template)
		// Your main service logic here
		logger.Info("Service stopped")
	},
}

func Execute(templates embed.FS) {
	template = templates
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
