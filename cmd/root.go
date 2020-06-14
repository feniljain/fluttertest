package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "fluttertest",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

//Execute acts as a entry point for CLI
func Execute() {
	rootCmd.Execute()
}
