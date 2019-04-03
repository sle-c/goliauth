package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var db string

func migrateCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "migrate",
		Short: "Create an `apps` table that store the key pairs",
		Long:  "Generate a table named `apps` in the database with the provided database url",
	}

	upCmd := &cobra.Command{
		Use:   "up",
		Short: "Execute create table query",
		Long:  "Execute create table query with provided database url",
		Run:   migrateApp,
	}

	downCmd := &cobra.Command{
		Use:   "down",
		Short: "Delete apps table",
		Long:  "Delete apps table with provided database url",
		Run:   rollbackApp,
	}

	upCmd.Flags().StringVarP(
		&db,
		"db",
		"d",
		"",
		"Url to connect to your database",
	)

	downCmd.Flags().StringVarP(
		&db,
		"db",
		"d",
		"",
		"Url to connect to your database",
	)

	command.AddCommand(
		upCmd,
		downCmd,
	)

	return command
}

func migrateApp(cmd *cobra.Command, args []string) {
	fmt.Println(db)
}

func rollbackApp(cmd *cobra.Command, args []string) {
	fmt.Println(db)
}
