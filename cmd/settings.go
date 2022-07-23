package cmd

// dpsctl default settings
const (
	LoginClientId		 						 = "w47MtsmRGb3DDTKqWdXMy9L7KudD5nDq"
	LoginScope                   = "openid offline_access profile email"
	LoginAudience                = ""

	DeviceCodeUrl								 = "https://twdpsio.us.auth0.com/oauth/device/code"
	AuthenticationUrl            = "https://twdpsio.us.auth0.com/oauth/token"

	ConfigEnvDefault             = "DPSCTL"
	ConfigFileDefaultName        = "config"
	ConfigFileDefaultType        = "yaml"
	ConfigFileDefaultLocation    = "/.dpsctl" // path will begin with $HOME dir
	ConfigFileDefaultLocationMsg = "config file (default is $HOME/.dpsctl/config.yaml)"
)