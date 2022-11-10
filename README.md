<div align="center">
	<p>
		<img alt="Thoughtworks Logo" src="https://raw.githubusercontent.com/ThoughtWorks-DPS/static/master/thoughtworks_flamingo_wave.png?sanitize=true" width=200 />
    <br />
		<img alt="DPS Title" src="https://raw.githubusercontent.com/ThoughtWorks-DPS/static/master/EMPCPlatformStarterKitsImage.png?sanitize=true" width=350/>
	</p>
  <br />
  <h3>dpsctl</h3>
    <a href="https://app.circleci.com/pipelines/github/ThoughtWorks-DPS/dpsctl"><img src="https://circleci.com/gh/ThoughtWorks-DPS/dpsctl.svg?style=shield"></a> <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
</div>
<br />

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/ThoughtWorks-DPS/dpsctl)

The cli is the primary tool for developments teams using platform starter kit based infrastructure sources.  

1. [Walkthrough](doc/auth0-device-auth-flow.md) of device-authorization-flow
2. [Quickstart](doc/quickstart.md)

### development

#### Config

You must first generate a config. Ensure your local one password CLI is logged in. Then you can run `op inject -i tpl/config.go.tpl -o cmd/config.go`

By default the cli is configured to attempt to authenticate against a production auth0 tenant ('twdpsio').  

In order to test against the dev-twdpsio auth0 tenant set the following values to the dev-tenant settings.  
```bash
export DPSCTL_IDPISSUERURL=https://dev-twdpsio.us.auth0.com/
export DPSCTL_LOGINCLIENTID=<insert dev-twdpsio dev-dpsctl application client id>
```

Clear your ~/.dpsctl/config.yaml file and re-login.  

To obtain non default kubeconfig, include --cluster flag with `dpsctl get kubeconfig` command.  
