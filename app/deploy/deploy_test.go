package deploy

import (
	"deploy/app/conf"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/mock"
	"testing"
)
var (
	_  = conf.Init()
	mk = new(ec2Mk)
	d  = Deployer{
		ec2: mk,
	}
)

func TestDeployer(t *testing.T) {
	// TODO: complete the test
}

type ec2Mk struct {
	mock.Mock
	waitCh            chan bool
}

func (e *ec2Mk) GetByAmi(ami string) []*ec2.Instance {
	return nil
}

func (e *ec2Mk) RunInstance(ami string) *ec2.Reservation {
	return nil
}

func (e *ec2Mk) WaitInstanceUp(instanceId *string) {
	<- e.waitCh
}

func (e *ec2Mk) TerminateInstance(instanceId *string) {

}

