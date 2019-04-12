package cmd

import (
	"fmt"

	"github.com/omnisyle/goliauth/goliauth/internal/app"
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
		Short: "Create an app with a public/secret key pair",
		Args:  cobra.MinimumNArgs(1),
		Run:   createApp,
	}

	getCmd := &cobra.Command{
		Use:   "get [public key]",
		Short: "Get an app using public key",
		Args:  cobra.MinimumNArgs(1),
		Run:   getApp,
	}

	delCmd := &cobra.Command{
		Use:   "delete [public key]",
		Short: "Delete an app using public key",
		Args:  cobra.MinimumNArgs(1),
		Run:   deleteApp,
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
		getCmd,
		delCmd,
	)

	return command
}

func createApp(cmd *cobra.Command, args []string) {
	appName := args[0]
	keyPair := app.CreateApp(appName, db)
	printResult(keyPair)
}

func getApp(cmd *cobra.Command, args []string) {
	publicKey := args[0]
	keyPair := app.GetApp(publicKey, db)
	printResult(keyPair)
}

func deleteApp(cmd *cobra.Command, args []string) {
	publicKey := args[0]
	result := app.DeleteApp(publicKey, db)

	if result {
		fmt.Println("Deleted")
	} else {
		fmt.Println("Failed to delete")
	}
}

func printResult(keyPair *app.App) {
	fmt.Println("Name: ", keyPair.Name)
	fmt.Println("Public Key: ", keyPair.PublicKey)
	fmt.Println("Secret Key: ", keyPair.SecretKey)
}
