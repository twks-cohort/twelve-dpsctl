package cmd

import (
	"fmt"
	//"runtime"
	//"os/exec"
	//"dpsctl/clients"
	//"dpsctl/clients/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:               "config",
	Short:             "Get Config",
	Long:              `Write configt`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config command")
		fmt.Printf("%s\n", viper.GetString("DeviceCodeUrl"))
		//viper.SetDefault("DeviceCodeUrl", DeviceCodeUrl)
	},
}

func init() {
	getCmd.AddCommand(configCmd)
}
