package cmd

import (
	"fmt"
	//"runtime"
	//"os/exec"
	//"dpsctl/clients"
	//"dpsctl/clients/models"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:               "get",
	Short:             "Get dpsctl configuration values",
	Long:              `Write dpsctl configuration values to stdout`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get command requires a valid argument")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
