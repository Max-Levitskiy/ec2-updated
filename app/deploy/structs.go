package deploy

import "deploy/app/aws"

type Deployer struct {
	ec2 aws.Ec2
}

type AppState struct {
	ec2 []struct{
		id string
		ami string
	}
}
