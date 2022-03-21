package cmd

// dpsctl default settings
const (
	LoginClientId		 						 = "2EOaJoAdFLSs1fxwi3lBc1ad1P0k5pPG"
	LoginScope                   = "openid offline_access profile email"
	LoginAudience                = ""

	DeviceCodeUrl								 = "https://dev-twdpsio.us.auth0.com/oauth/device/code"
	AuthenticationUrl            = "https://dev-twdpsio.us.auth0.com/oauth/token"

	ConfigEnvDefault             = "DPSCTL"
	ConfigFileDefaultName        = "config"
	ConfigFileDefaultType        = "yaml"
	ConfigFileDefaultLocation    = "/.dpsctl" // path will begin with $HOME dir
	ConfigFileDefaultLocationMsg = "config file (default is $HOME/.dpsctl/config.yaml)"
)