package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var db string

func appCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "app",
		Short: "Create a public/secret key pair as an app with a name",
	}

	createCmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a app with a public/secret key pair",
		Args:  cobra.MinimumNArgs(1),
		Run:   createApp,
	}

	command.PersistentFlags().StringVarP(
		&db,
		"db",
		"d",
		"",
		"Url to connect to the database",
	)

	command.MarkFlagRequired("db")

	command.AddCommand(
		createCmd,
	)

	return command
}

func createApp(cmd *cobra.Command, args []string) {
	name := args[0]
	fmt.Println(name)
	fmt.Println(db)
}
