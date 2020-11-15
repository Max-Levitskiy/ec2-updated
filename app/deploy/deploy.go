package deploy

import (
	"deploy/app/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"
	"sync"
)

func NewDeployer() *Deployer {
	d := &Deployer{}
	d.ec2 = aws.NewEc2Impl()
	return d
}

func (d *Deployer) DeployApp(oldAmi string, newAmi string) {
	required := viper.GetInt("app.ec2.required-amount")
	var wg sync.WaitGroup
	wg.Add(required)
	fmt.Printf("Running %d new instances\n", required, )
	for i := 0; i < required; i++ {
		instance := d.ec2.RunInstance(newAmi)
		go func(e ec2.Reservation) {
			id := e.Instances[0].InstanceId
			fmt.Printf("Waiting instance %s to be started\n", *id)
			d.ec2.WaitInstanceUp(id)
			wg.Done()
			fmt.Printf("Instance %s started\n", *id)
		}(*instance)
	}
	wg.Wait()

	fmt.Printf("Check and terminate old instances with AMI %s\n", oldAmi)
	oldInstances := d.ec2.GetByAmi(oldAmi)
	i := len(oldInstances)
	if i > 0 {
		fmt.Printf("Found %d instances\n", i)
		for _, instance := range oldInstances {
			id := instance.InstanceId
			fmt.Printf("Terminating instance %s\n", *id)
			d.ec2.TerminateInstance(id)
		}
	} else {
		fmt.Println("No old instances found")
	}
	fmt.Println("Done")
}

