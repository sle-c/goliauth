package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goliauth",
	Short: "A mini CLI to generate random key and public/secret key pairs for API",
}

// Execute chains all the commands together under the root command which is goliauth
// Usage: `goliauth [command]`
func Execute() {
	rootCmd.AddCommand(
		keyCmd(),
		migrateCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
