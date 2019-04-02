package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goliauth",
	Short: "A mini CLI to generate random key and credentials for API",
	Long:  "A mini CLI with helper methods to create credentials for your API",
}

func Execute() {
	rootCmd.AddCommand(
		keyCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
