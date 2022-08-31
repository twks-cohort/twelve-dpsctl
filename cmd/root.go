package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Version = "snapshot"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dpsctl",
	Short: "empc platform starter kit cli",
	Long:  `cli for use with empc lab platform.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitOnError(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// you may specify the config file and location. Viper supports the following file types based on extension:
	// JSON, TOML, YAML, HCL, INI, envfile and Java Properties files
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", ConfigFileDefaultLocationMsg)
}

// initConfig sets the config values based on the following order of precedent:
// ENV variables
// Config file definitions
// Default values from settings.go
func initConfig() {

	viper.SetDefault("LoginClientID", LoginClientId)
	viper.SetDefault("LoginScope", LoginScope)
	viper.SetDefault("LoginAudience", LoginAudience)
	viper.SetDefault("IdpIssuerUrl", IdpIssuerUrl)
	viper.SetDefault("DefaultCluster", DefaultCluster)
	viper.SetDefault("TeamsApi", TeamsApi)

	viper.SetEnvPrefix(ConfigEnvDefault)
	viper.AutomaticEnv()

	if cfgFile != "" {
		// Use config file from the flag if specified.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(defaultConfigLocation())
		viper.SetConfigName(ConfigFileDefaultName)
	}

	// If a config file is found, read it in, else write a blank.
	if err := viper.ReadInConfig(); err != nil {
		configFileLocation := defaultConfigLocation()
		configFilePath := configFileLocation + "/" + ConfigFileDefaultName + "." + ConfigFileDefaultType

		exitOnError(os.MkdirAll(configFileLocation, 0700))
		fmt.Println("created " + configFilePath)
		emptyFile, err := os.Create(configFilePath)
		exitOnError(err)
		emptyFile.Close()
	}
	viper.WriteConfig() //nolint:errcheck
}

func defaultConfigLocation() string {
	home, err := homedir.Dir()
	exitOnError(err)
	return home + ConfigFileDefaultLocation
}
