package cmd

import (
	"dpsctl/clients"
	"dpsctl/clients/models"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

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
	submitCode(deviceCode.VerificationUriComplete)

	clients.Authenticate(deviceCode)
}

func checkForWSL() (bool, error) {
	dat, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return false, err
	}

	if strings.Contains(string(dat), "microsoft") {
		return true, nil
	}

	return false, nil
}

func submitHandler(url string) error {
	fmt.Println("dpsctl will attempt to open a browser window where you can authenticate and verify your laptop.")
	fmt.Println("If the window does not open, go to the link below.") //nolint:govet
	fmt.Printf("%s\n", url)

	switch runtime.GOOS {
	case "linux":
		err := exec.Command("xdg-open", url).Start()

		//If this failed, we might be running on WSL in Windows
		//Check if that is the case and launch a different command
		if err != nil {
			isWsl, err := checkForWSL()
			if err != nil {
				// Problem parsing /proc file, could be permission or other issue
				return transientError{err: err}
			}

			if isWsl {
				return exec.Command("sensible-browser", url).Start()
			}

			return err
		}

		return nil
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}

func submitCode(url string) {
	err := submitHandler(url)
	if err != nil {
		terr := transientError{}
		if errors.As(err, &terr) {
			log.Printf("There was a problem detecting the underlying OS \n")
			log.Fatal(err.Error())
		} else {
			exitOnError(err)
		}
	}
}
