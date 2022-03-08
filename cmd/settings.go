package cmd

// dpsctl default settings
const (
	LoginClientId		 						 = "B4jm7Wv4fjOEPqg1gjXIUUxEa6eg1HvB"
	LoginScope                   = "openid offline_access email"
	LoginAudience                = "https://mapi.twdps.digital/v1"

	DeviceCodeUrl								 = "https://dev-twdpsio.us.auth0.com/oauth/device/code"

	ConfigEnvDefault             = "DPSCTL"
	ConfigFileDefaultName        = "config"
	ConfigFileDefaultType        = "yaml"
	ConfigFileDefaultLocation    = "/.dpsctl" // path will begin with $HOME dir
	ConfigFileDefaultLocationMsg = "config file (default is $HOME/.dpsctl/config.yaml)"
)