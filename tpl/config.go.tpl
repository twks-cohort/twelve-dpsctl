package cmd

type ClusterConfig struct {
clusterName string
clusterEndpoint string
base64CertificateAuthorityData string
}

var cluster1 = ClusterConfig{
clusterName: "sandbox-ap-southeast-2",
clusterEndpoint: "{{ op://cohorts/twelve-platform-sandbox-ap-southeast-2/cluster-endpoint }}",
base64CertificateAuthorityData: "{{ op://cohorts/twelve-platform-sandbox-ap-southeast-2/base64-certificate-authority-data }}",
}



var clusters = []ClusterConfig{ cluster1 }

const (
LoginClientId		 						    = "{{ op://cohorts/team-twelve-svc-auth0/dev-8zg3kpi25tnc00ds-dev-dpsctl-client-id}}"
LoginScope                      = "openid offline_access profile email"
LoginAudience                   = ""

IdpIssuerUrl								    = "https://dev-8zg3kpi25tnc00ds.us.auth0.com/"

ConfigEnvDefault                = "DPSCTL"
ConfigFileDefaultName           = "config"
ConfigFileDefaultType           = "yaml"
ConfigFileDefaultLocation       = "/.dpsctl" // path will begin with $HOME dir
ConfigFileDefaultLocationMsg    = "config file (default is $HOME/.dpsctl/config.yaml)"

DefaultCluster                  = "sandbox-ap-southeast-2"
TeamsApi                        = "http://localhost:8000"
)
