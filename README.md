# dpsctl

Example delivery platform cli

### development



By default the cli is configured to attempt to authenticate against a production auth0 tenant ('twdpsio').  

In order to test against the dev-twdpsio auth0 tenant set the following values to the dev-tenant settings.  
```bash
export DPSCTL_DEVICECODEURL=https://dev-twdpsio.us.auth0.com/oauth/device/code
export DPSCTL_AUTHENTICATIONURL=https://dev-twdpsio.us.auth0.com/oauth/token
export DPSCTL_LOGINCLIENTID=<insert dev-twdpsio dev-dpsctl application client id>
```
