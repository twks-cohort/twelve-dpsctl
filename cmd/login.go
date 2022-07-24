package cmd

import (
	"fmt"
	"runtime"
	"os/exec"
	"dpsctl/clients"
	"dpsctl/clients/models"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:               "login",
	Short:             "Login to empc platform starter kit lab",
	Long:              `Login to empc platform starter kit lab using authenticated Github credentials`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		login(clients.RequestDeviceCode())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login(deviceCode models.DeviceCode) {
	// provide link for browser based authentication and device verfication
	// and attempt to automatically open a browser window for the user
	submitcode(deviceCode.VerificationUriComplete)

	clients.Authenticate(deviceCode)
}


func submitcode(url string) {
	var err error

	fmt.Println("dpsctl will attempt to open a browser window where you can authenticate and verify your laptop.")
  fmt.Println("If the window does not open, go to the link below.\n")
  fmt.Printf("%s\n", url)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	exitOnError(err)
}