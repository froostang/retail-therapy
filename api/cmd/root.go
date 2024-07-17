package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "my-service",
	Short: "A Cobra-based service daemon with Zap logging options",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		// TODO Add HTTP Server

		// Your main service logic here
		logger.Info("Service started")

		// Example of logging an error
		logger.Error("An error occurred", zap.Error(fmt.Errorf("example error")))

		// Example of logging a warning
		logger.Warn("A warning message")

		// Example of logging an info message
		logger.Info("An info message")

		// Example of logging a debug message
		logger.Debug("A debug message")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
