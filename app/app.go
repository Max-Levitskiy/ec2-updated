package app

import (
	"deploy/app/conf"
	"deploy/app/deploy"
)

func Run(oldAmi string, newAmi string) {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	deploy.NewDeployer().DeployApp(oldAmi, newAmi)
}

