### Quickstart

Login to generate local access credentials. You will be re-directed to authenticate via GitHub.  
```
$ dpsctl login
```
This will create a configuration file at ~/.dpsctl/confnig.yaml  

Among the credentials generated will be a WJT bearer and refresh token that is used by a kubernetes oicd provider to authenticate your access to the kubernetes api. The token contains your claims in the form of your team memberships within the authorizing github Organization. You will only be able to access the kubernetes api where the oidc provider can both successfully authenticate your token, and where at least one prior clusterroldbindings matches your claims.  

Generate a kubeconfig.  
```bash
$ dpsctl get kubeconfig > mykubeconfig
```

The resulting contents of the file will look something like this:
```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ABCDefgh12345==
    server: https://12341567890.gr7.us-east-1.eks.amazonaws.com
  name: prod-us-east-1
contexts:
- context:
    cluster: prod-us-east-1
    user: oidc-user@prod-us-east-1
  name: prod-us-east-1
current-context: prod-us-east-1
kind: Config
preferences: {}
users:
- name: oidc-user@prod-us-east-1
  user:
    auth-provider:
      config:
        client-id: ABCDefgh12345
        idp-issuer-url: https://twdpsio.us.auth0.com/
        refresh-token: ABCDefgh12345
      name: oidc
```
By default the empc platform start kit lab provides credentials for a cluster named prod-us-east-1, since that is where the example developer environments exist.  
