package cmd

import (
	"fmt"

	"github.com/omnisyle/goliauth"
	"github.com/spf13/cobra"
)

func keyCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "key",
		Short: "Generate a 32 bit random key",
		Long:  "Generate a 32 bit random key",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%x", goliauth.NewRandomKey())
		},
	}

	return command
}
