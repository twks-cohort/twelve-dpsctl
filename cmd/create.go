package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create dpsctl objects",
	Long:              `Write dpsctl objects to dynamodb`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create command requires a valid argument")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
