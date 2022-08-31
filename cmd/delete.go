package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete dpsctl objects",
	Long:              `Delete dpsctl objects, such as teams and members`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete command requires a valid argument")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
