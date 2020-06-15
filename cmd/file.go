package cmd

import (
	"flutterTest/parser"
	"fmt"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Pass the required fileName as the argument",
	Run: func(cm *cobra.Command, args []string) {
		fileName := args[0]
		err := parser.Parse(fileName)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Some error occurred!")
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
