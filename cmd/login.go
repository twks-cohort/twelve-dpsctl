package cmd

import (
	"fmt"
	"dpsctl/clients"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:               "login",
	Short:             "Login to twdps lab",
	Long:              `Login to twdps lab using authenticated Github credentials`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login command")

		clients.RequestDeviceCode()
		// fmt.Println(deviceCode)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}